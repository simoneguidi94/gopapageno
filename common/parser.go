import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

/*
parsingStats contains some statistics about the parse.
*/
type parsingStats struct {
	StackPoolSize      int
	StackPtrPoolSize   int
	AllocMemTime       time.Duration
	LexTime            time.Duration
	NumTokens          int
	ParseTime          time.Duration
	RemainingStacks    int
	RemainingStackPtrs int
}

/*
Stats contains some statistics that may be checked after a call to ParseString or ParseFile
*/
var Stats parsingStats

/*
threadContext contains the data required by each thread, basically the thread number,
an input list, a parsing stack and a list that contains all the generated nonterminals.
*/
type threadContext struct {
	num             int
	input           *listOfStacks
	newNonTerminals *listOfStacks
	stack           *listOfStackPtrs
	result          string
}

/*
threadJob is the parsing function executed in parallel by each thread.
It takes as input a threadContext and a channel where it eventually sends the result.
*/
func threadJob(context threadContext, c chan threadContext) {
	threadNum := context.num
	input := context.input
	inputIterator := input.HeadIterator()
	newNonTerminalsList := context.newNonTerminals
	stack := context.stack

	numYieldsPrec := 0

	var pos int
	var sym *symbol
	var ruleNum uint16
	var lhs uint16
	var lhsSym *symbol
	var rhs []uint16
	var rhsSymbols []*symbol
	rhsBuf := make([]uint16, _MAX_RHS_LEN)
	rhsSymbolsBuf := make([]*symbol, _MAX_RHS_LEN)

	tokensRead := 0

	newNonTerm := &symbol{0, _NO_PREC, nil, nil, nil}

	//Get the first symbol from the input list
	inputSym := inputIterator.Next()

	//Except for the first thread, ignore the first symbol and get the second one.
	if threadNum > 0 {
		inputSym = inputIterator.Next()
		tokensRead++
	}

	//Iterate over all the input list
	for tokensRead <= input.Length()-1 {
		//stack.Println()

		//If the current token is a nonterminal, push it onto the stack with no precedence relation
		if !isTerminal(inputSym.Token) {
			//fmt.Printf("Pushed (%s, %s)\n", TokenToString(inputSym.Token), precToString(NO_PREC))
			inputSym.Precedence = _NO_PREC
			stack.Push(inputSym)
			inputSym = inputIterator.Next()
			tokensRead++
			continue
		}

		//Find the first terminal on the stack and get the precedence between it and the current input token
		firstTerminal := stack.FirstTerminal()
		prec := getPrecedence(firstTerminal.Token, inputSym.Token)

		switch prec {
		//If it yields precedence, push the input token onto the stack with that precedence relation.
		//Also increment the counter of the number of tokens yielding precedence.
		case _YIELDS_PREC:
			//fmt.Printf("Pushed (%s, %s)\n", TokenToString(inputSym.Token), precToString(prec))
			inputSym.Precedence = _YIELDS_PREC
			stack.Push(inputSym)
			numYieldsPrec++
			inputSym = inputIterator.Next()
			tokensRead++
		//If it's equal in precedence, push the input token onto the stack with that precedence relation
		case _EQ_PREC:
			//fmt.Printf("Pushed (%s, %s)\n", TokenToString(inputSym.Token), precToString(prec))
			inputSym.Precedence = _EQ_PREC
			stack.Push(inputSym)
			inputSym = inputIterator.Next()
			tokensRead++
		//If it takes precedence, the next action depends on whether there are tokens that yield precedence onto the stack.
		case _TAKES_PREC:
			//If there are no tokens yielding precedence on the stack, push the input token onto the stack
			//with take precedence as precedence relation
			if numYieldsPrec == 0 {
				//fmt.Printf("Pushed (%s, %s)\n", TokenToString(inputSym.Token), precToString(prec))
				inputSym.Precedence = _TAKES_PREC
				stack.Push(inputSym)
				inputSym = inputIterator.Next()
				tokensRead++
				//Otherwise, perform a reduction
			} else {
				pos = _MAX_RHS_LEN - 1

				//Pop tokens from the stack until one that yields precedence is reached, saving them in rhsBuf
				sym = stack.Pop()

				for sym.Precedence != _YIELDS_PREC {
					rhsSymbolsBuf[pos] = sym
					rhsBuf[pos] = sym.Token
					sym = stack.Pop()
					pos--
				}
				rhsSymbolsBuf[pos] = sym
				rhsBuf[pos] = sym.Token

				//Pop one last token, if it's a nonterminal add it to rhsBuf, otherwise ignore it (push it again onto the stack)
				sym = stack.Pop()

				if isTerminal(sym.Token) {
					stack.Push(sym)
				} else {
					pos--
					rhsSymbolsBuf[pos] = sym
					rhsBuf[pos] = sym.Token
					stack.UpdateFirstTerminal()
				}

				//Obtain the actual rhs from the buffers
				rhsSymbols = rhsSymbolsBuf[pos:]
				rhs = rhsBuf[pos:]

				//Find corresponding lhs and ruleNum
				lhs, ruleNum = findMatch(rhs)

				//If a rule with that rhs does not exist, abort the parsing
				if lhs == _EMPTY {
					/*fmt.Print("Error, could not find a reduction for ")
					for i := 0; i < len(rhs); i++ {
						fmt.Print(TokenToString(rhs[i]))
						fmt.Print(",")
					}
					fmt.Println()*/

					context.result = "failure"
					c <- context

					return
				}

				/*fmt.Print("Reduced ")
				for i := 0; i < len(rhs); i++ {
					fmt.Print(TokenToString(rhs[i]))
					fmt.Print(",")
				}
				fmt.Print(" -> ")
				fmt.Println(TokenToString(lhs))*/

				//Push the new nonterminal onto the appropriate list to save it
				newNonTerm.Token = lhs
				lhsSym = newNonTerminalsList.Push(newNonTerm)

				//Execute the semantic action
				function(threadNum, ruleNum, lhsSym, rhsSymbols)

				//Push the new nonterminal onto the stack
				stack.Push(lhsSym)

				//Decrement the counter of the number of tokens yielding precedence
				numYieldsPrec--
			}
		//If there's no precedence relation, abort the parsing
		case _NO_PREC:
			//fmt.Printf("Error, no precedence relation between %s and %s\n", TokenToString(firstTerminal.Token), TokenToString(inputSym.Token))

			context.result = "failure"
			c <- context

			return
		}
	}

	context.result = "success"
	c <- context
}

var cpuprofileFile *os.File = nil

func SetCPUProfileFile(file *os.File) {
	cpuprofileFile = file
}

/*
ParseString parses a string in parallel using an operator precedence grammar.
It takes as input a string as a slice of bytes and the number of threads, and returns a boolean
representing the success or failure of the parsing and the symbol at the root of the syntactic tree (if successful).
*/
func ParseString(str []byte, numThreads int) (bool, *symbol) {
	rawInputSize := len(str)

	avgCharsPerToken := float64(2)

	//The last multiplication by  is to account for the generated nonterminals
	stackPoolSize := int(math.Ceil((float64(rawInputSize)/avgCharsPerToken)/float64(_STACK_SIZE))) * 2
	stackPtrPoolSize := int(math.Ceil((float64(rawInputSize) / avgCharsPerToken) / float64(_STACK_PTR_SIZE)))

	Stats.StackPoolSize = stackPoolSize
	Stats.StackPtrPoolSize = stackPtrPoolSize

	//Alloc memory required by both the lexer and the parser.
	//The call to runtime.GC() avoids the garbage collector to run concurrently with the parser
	start := time.Now()

	stackPool := newStackPool(stackPoolSize)

	stackPtrPool := newStackPtrPool(stackPtrPoolSize)

	lexerPreallocMem(rawInputSize, numThreads)

	parserPreallocMem(rawInputSize, numThreads)

	runtime.GC()

	Stats.AllocMemTime = time.Since(start)

	//Lex the file to obtain the input list
	start = time.Now()

	input, err := lex(str, stackPool)

	Stats.LexTime = time.Since(start)

	//If lexing fails, abort the parsing
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}

	if cpuprofileFile != nil {
		if err := pprof.StartCPUProfile(cpuprofileFile); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	Stats.NumTokens = input.Length()

	if input.Length() == 0 {
		return true, nil
	}

	start = time.Now()

	var result *symbol = nil

	//If there are not enough stacks in the input, reduce the number of threads.
	//This is because the input is split by splitting stacks, not stack contents.
	if input.NumStacks() < numThreads {
		fmt.Println("There are less stacks than threads, reducing the number of threads to", input.NumStacks())
		numThreads = input.NumStacks()
	}

	//Split the input list
	inputLists := input.Split(numThreads)

	newNonTerminalsLists := make([]listOfStacks, numThreads)

	for i := 0; i < numThreads; i++ {
		newNonTerminalsLists[i] = newLos(stackPool)
	}

	threadContexts := make([]threadContext, numThreads)

	c := make(chan threadContext)

	//Create the thread contexts and run the threads
	for i := 0; i < numThreads; i++ {
		curThreadStack := newLosPtr(stackPtrPool)
		//If the thread is the first, push a # onto the stack
		if i == 0 {
			curThreadStack.Push(&symbol{_TERM, _NO_PREC, nil, nil, nil})
			//Otherwise, push the first token onto the stack
		} else {
			curInputListIter := inputLists[i].HeadIterator()
			sym := curInputListIter.Next()
			sym.Precedence = _NO_PREC
			curThreadStack.Push(sym)
		}
		//If the thread is the last, push a # onto the input list
		if i == numThreads-1 {
			inputLists[i].Push(&symbol{_TERM, _NO_PREC, nil, nil, nil})
			//Otherwise, push onto the input list the first token of the next input list
		} else {
			nextInputListIter := inputLists[i+1].HeadIterator()
			inputLists[i].Push(nextInputListIter.Next())
		}
		threadContexts[i] = threadContext{i, &inputLists[i], &newNonTerminalsLists[i], &curThreadStack, ""}

		//fmt.Print("Thread", i, " input: ")
		//threadContexts[i].input.Println()
		//fmt.Print("Thread", i, " stack: ")
		//threadContexts[i].stack.Println()

		go threadJob(threadContexts[i], c)

		/*threadContexts[i] = <-c

		fmt.Println("Thread", threadContexts[i].num, "finished parsing")
		fmt.Println("Result:", threadContexts[i].result)
		fmt.Print("Partial stack: ")
		threadContexts[i].stack.Println()

		if threadContexts[i].result == "failure" {
			fmt.Printf("Time to parse it: %s\n", time.Since(start))
			return false
		}*/
	}

	//Wait for each thread to finish its job
	for i := 0; i < numThreads; i++ {
		threadContext := <-c

		threadContexts[threadContext.num] = threadContext

		//fmt.Println("Thread", threadContexts[i].num, "finished parsing")
		//fmt.Println("Result:", threadContexts[i].result)
		//fmt.Print("Partial stack: ")
		//threadContexts[i].stack.Println()

		//If one of the threads fails, abort the parsing
		if threadContexts[i].result == "failure" {
			Stats.ParseTime = time.Since(start)
			return false, nil
		}
	}

	//If the number of threads is greater than one, a final pass is required
	if numThreads > 1 {
		//Create the final input by joining together the stacks from the previous step
		finalPassInput := newLos(stackPool)
		for i := 0; i < numThreads; i++ {
			iterator := threadContexts[i].stack.HeadIterator()
			//Ignore the first token
			iterator.Next()
			sym := iterator.Next()
			for sym != nil {
				finalPassInput.Push(sym)
				sym = iterator.Next()
			}
		}

		finalPassNewNonTerminalsList := newLos(stackPool)

		finalPassStack := newLosPtr(stackPtrPool)
		finalPassStack.Push(&symbol{_TERM, _NO_PREC, nil, nil, nil})
		finalPassThreadContext := threadContext{0, &finalPassInput, &finalPassNewNonTerminalsList, &finalPassStack, ""}

		//fmt.Print("Final pass thread input: ")
		//finalPassThreadContext.input.Println()
		//fmt.Print("Final pass thread stack: ")
		//finalPassThreadContext.stack.Println()

		go threadJob(finalPassThreadContext, c)

		finalPassThreadContext = <-c

		//fmt.Println("Final thread finished parsing")
		//fmt.Println("Result:", finalPassThreadContext.result)
		//fmt.Print("Final stack: ")
		//finalPassThreadContext.stack.Println()

		if finalPassThreadContext.result == "failure" {
			Stats.ParseTime = time.Since(start)
			return false, nil
		}

		//Pop tokens from the stack until a nonterminal is found
		sym := finalPassThreadContext.stack.Pop()

		for isTerminal(sym.Token) {
			sym = finalPassThreadContext.stack.Pop()
		}

		//sym.PrintTreeln()

		//Set the result as the nonterminal symbol
		result = sym
	} else {
		//Pop tokens from the stack until a nonterminal is found
		sym := threadContexts[0].stack.Pop()

		for isTerminal(sym.Token) {
			sym = threadContexts[0].stack.Pop()
		}

		//sym.PrintTreeln()

		//Set the result as the nonterminal symbol
		result = sym
	}

	Stats.ParseTime = time.Since(start)

	Stats.RemainingStacks = stackPool.Remainder()
	Stats.RemainingStackPtrs = stackPtrPool.Remainder()

	return true, result
}

/*
ParseFile parses a file in parallel using an operator precedence grammar.
It takes as input a filename and the number of threads, and returns a boolean
representing the success or failure of the parsing and the symbol at the root of the syntactic tree (if successful).
*/
func ParseFile(filename string, numThreads int) (bool, *symbol) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}

	return ParseString(bytes, numThreads)
}
