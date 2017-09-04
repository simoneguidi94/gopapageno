package arithmetic

import (
	"errors"
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
	NumLexThreads                           int
	NumParseThreads                         int
	StackPoolSizes                          []int
	StackPoolNewNonterminalsSizes           []int
	StackPtrPoolSizes                       []int
	StackPoolSizeFinalPass                  int
	StackPoolNewNonterminalsSizeFinalPass   int
	StackPtrPoolSizeFinalPass               int
	AllocMemTime                            time.Duration
	CutPoints                               []int
	LexTimes                                []time.Duration
	LexTimeTotal                            time.Duration
	NumTokens                               []int
	NumTokensTotal                          int
	ParseTimes                              []time.Duration
	RecombiningStacksTime                   time.Duration
	ParseTimeFinalPass                      time.Duration
	ParseTimeTotal                          time.Duration
	RemainingStacks                         []int
	RemainingStacksNewNonterminals          []int
	RemainingStackPtrs                      []int
	RemainingStacksFinalPass                int
	RemainingStacksNewNonterminalsFinalPass int
	RemainingStackPtrsFinalPass             int
}

/*
Stats contains some statistics that may be checked after a call to ParseString or ParseFile
*/
var Stats parsingStats

/*
threadContext contains the data required by each thread, basically the thread number,
an input list, a parsing stack and a list that contains all the generated nonterminals.
*/
type parseResult struct {
	threadNum int
	stack     *listOfStackPtrs
	success   bool
}

type lexResult struct {
	threadNum int
	tokenList *listOfStacks
	success   bool
}

/*
threadJob is the parsing function executed in parallel by each thread.
It takes as input a threadContext and a channel where it eventually sends the result.
*/
func threadJob(numThreads int, threadNum int, finalPass bool, input *listOfStacks, nextSym *symbol, stackPool *stackPool, stackPtrPool *stackPtrPool, c chan parseResult) {
	start := time.Now()

	inputIterator := input.HeadIterator()
	newNonTerminalsList := newLos(stackPool)
	stack := newLosPtr(stackPtrPool)
	//tokensRead := 0

	//If the thread is the first, push a # onto the stack
	if threadNum == 0 {
		stack.Push(&symbol{_TERM, _NO_PREC, nil, nil, nil})
		//Otherwise, push the first token onto the stack
	} else {
		sym := inputIterator.Next()
		sym.Precedence = _NO_PREC
		stack.Push(sym)
		//tokensRead++
	}
	//If the thread is the last, push a # onto the input list
	if threadNum == numThreads-1 {
		input.Push(&symbol{_TERM, _NO_PREC, nil, nil, nil})
		//Otherwise, push onto the input list the first token of the next input list
	} else {
		input.Push(nextSym)
	}

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

	newNonTerm := &symbol{0, _NO_PREC, nil, nil, nil}

	//Get the first symbol from the input list
	inputSym := inputIterator.Next()

	//Iterate over all the input list
	for inputSym != nil {
		//stack.Println()

		//If the current token is a nonterminal, push it onto the stack with no precedence relation
		if !isTerminal(inputSym.Token) {
			//fmt.Printf("Pushed (%s, %s)\n", TokenToString(inputSym.Token), precToString(NO_PREC))
			inputSym.Precedence = _NO_PREC
			stack.Push(inputSym)
			inputSym = inputIterator.Next()
			//tokensRead++
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
			//tokensRead++
		//If it's equal in precedence, push the input token onto the stack with that precedence relation
		case _EQ_PREC:
			//fmt.Printf("Pushed (%s, %s)\n", TokenToString(inputSym.Token), precToString(prec))
			inputSym.Precedence = _EQ_PREC
			stack.Push(inputSym)
			inputSym = inputIterator.Next()
			//tokensRead++
		//If it takes precedence, the next action depends on whether there are tokens that yield precedence onto the stack.
		case _TAKES_PREC:
			//If there are no tokens yielding precedence on the stack, push the input token onto the stack
			//with take precedence as precedence relation
			if numYieldsPrec == 0 {
				//fmt.Printf("Pushed (%s, %s)\n", TokenToString(inputSym.Token), precToString(prec))
				inputSym.Precedence = _TAKES_PREC
				stack.Push(inputSym)
				inputSym = inputIterator.Next()
				//tokensRead++
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

					c <- parseResult{threadNum, nil, true}

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

			c <- parseResult{threadNum, nil, true}

			return
		}
	}

	c <- parseResult{threadNum, &stack, true}

	if !finalPass {
		Stats.ParseTimes[threadNum] = time.Since(start)
	} else {
		Stats.ParseTimeFinalPass = time.Since(start)
	}
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
func ParseString(str []byte, numThreads int) (*symbol, error) {
	rawInputSize := len(str)

	avgCharsPerToken := float64(2)

	//The last multiplication by  is to account for the generated nonterminals
	stackPoolBaseSize := math.Ceil((((float64(rawInputSize) / avgCharsPerToken) / float64(_STACK_SIZE)) / float64(numThreads)))
	stackPtrPoolBaseSize := math.Ceil(((float64(rawInputSize) / avgCharsPerToken) / float64(_STACK_PTR_SIZE)) / float64(numThreads))

	//Stats.StackPoolSize = stackPoolSize
	//Stats.StackPtrPoolSize = stackPtrPoolSize

	//Alloc memory required by both the lexer and the parser.
	//The call to runtime.GC() avoids the garbage collector to run concurrently with the parser
	start := time.Now()

	stackPools := make([]*stackPool, numThreads)
	stackPoolsNewNonterminals := make([]*stackPool, numThreads)
	stackPtrPools := make([]*stackPtrPool, numThreads)

	Stats.StackPoolSizes = make([]int, numThreads)
	Stats.StackPoolNewNonterminalsSizes = make([]int, numThreads)
	Stats.StackPtrPoolSizes = make([]int, numThreads)

	for i := 0; i < numThreads; i++ {
		stackPools[i] = newStackPool(int(stackPoolBaseSize * 1.2))
		Stats.StackPoolSizes[i] = int(stackPoolBaseSize * 1.2)
		stackPoolsNewNonterminals[i] = newStackPool(int(stackPoolBaseSize))
		Stats.StackPoolNewNonterminalsSizes[i] = int(stackPoolBaseSize)
		stackPtrPools[i] = newStackPtrPool(int(stackPtrPoolBaseSize))
		Stats.StackPtrPoolSizes[i] = int(stackPtrPoolBaseSize)
	}

	var stackPoolFinalPass *stackPool
	var stackPoolNewNonterminalsFinalPass *stackPool
	var stackPtrPoolFinalPass *stackPtrPool
	if numThreads > 1 {
		stackPoolFinalPass = newStackPool(int(math.Ceil(stackPoolBaseSize * 0.1 * float64(numThreads))))
		Stats.StackPoolSizeFinalPass = int(math.Ceil(stackPoolBaseSize * 0.1 * float64(numThreads)))
		stackPoolNewNonterminalsFinalPass = newStackPool(int(math.Ceil(stackPoolBaseSize * 0.05 * float64(numThreads))))
		Stats.StackPoolNewNonterminalsSizeFinalPass = int(math.Ceil(stackPoolBaseSize * 0.05 * float64(numThreads)))
		stackPtrPoolFinalPass = newStackPtrPool(int(math.Ceil(stackPtrPoolBaseSize * 0.1)))
		Stats.StackPtrPoolSizeFinalPass = int(int(math.Ceil(stackPtrPoolBaseSize * 0.1)))
	}

	lexerPreallocMem(rawInputSize, numThreads)

	parserPreallocMem(rawInputSize, numThreads)

	runtime.GC()

	Stats.AllocMemTime = time.Since(start)

	//Lex the file to obtain the input list
	start = time.Now()

	cutPoints, numLexThreads := findCutPoints(str, numThreads)

	Stats.NumLexThreads = numLexThreads
	Stats.LexTimes = make([]time.Duration, numLexThreads)
	Stats.CutPoints = cutPoints

	if numLexThreads < numThreads {
		fmt.Printf("It was not possible to find cut points for all %d threads.\n", numThreads)
		fmt.Printf("The number of lexing threads was reduced to %d.\n", numLexThreads)
	}

	lexC := make(chan lexResult)

	for i := 0; i < numLexThreads; i++ {
		go lex(i, str[cutPoints[i]:cutPoints[i+1]], stackPools[i], lexC)
	}

	lexResults := make([]lexResult, numLexThreads)

	for i := 0; i < numLexThreads; i++ {
		curLexResult := <-lexC
		lexResults[curLexResult.threadNum] = curLexResult

		if !curLexResult.success {
			Stats.LexTimeTotal = time.Since(start)
			return nil, errors.New("Lexing error")
		}
	}

	input := lexResults[0].tokenList

	for i := 1; i < numLexThreads; i++ {
		input.Merge(*lexResults[i].tokenList)
	}

	//input, err := lex(str, stackPool, lexC)

	Stats.LexTimeTotal = time.Since(start)

	//If lexing fails, abort the parsing
	/*if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}*/

	if cpuprofileFile != nil {
		if err := pprof.StartCPUProfile(cpuprofileFile); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	Stats.NumTokensTotal = input.Length()

	if input.Length() == 0 {
		return nil, nil
	}

	start = time.Now()

	var result *symbol = nil

	//If there are not enough stacks in the input, reduce the number of threads.
	//This is because the input is split by splitting stacks, not stack contents.
	if input.NumStacks() < numThreads {
		fmt.Println("There are less stacks than threads, reducing the number of threads to", input.NumStacks())
		numThreads = input.NumStacks()
	}

	Stats.NumParseThreads = numThreads
	Stats.ParseTimes = make([]time.Duration, numThreads)

	//Split the input list
	inputLists := input.Split(numThreads)

	Stats.NumTokens = make([]int, numThreads)
	for i := 0; i < numThreads; i++ {
		Stats.NumTokens[i] = inputLists[i].Length()
	}

	parseResults := make([]parseResult, numThreads)

	c := make(chan parseResult)

	//Create the thread contexts and run the threads
	for i := 0; i < numThreads; i++ {
		//fmt.Print("Thread", i, " input: ")
		//threadContexts[i].input.Println()
		//fmt.Print("Thread", i, " stack: ")
		//threadContexts[i].stack.Println()

		var nextSym *symbol = nil

		if i < numThreads-1 {
			nextInputListIter := inputLists[i+1].HeadIterator()
			nextSym = nextInputListIter.Next()
		}

		go threadJob(numThreads, i, false, &inputLists[i], nextSym, stackPoolsNewNonterminals[i], stackPtrPools[i], c)

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
		curParseResults := <-c

		parseResults[curParseResults.threadNum] = curParseResults

		//fmt.Println("Thread", threadContext.num, "finished parsing")
		//fmt.Println("Result:", threadContext.result)
		//fmt.Print("Partial stack: ")
		//threadContext.stack.Println()

		//If one of the threads fails, abort the parsing
		if !curParseResults.success {
			Stats.ParseTimeTotal = time.Since(start)
			return nil, errors.New("Parsing error")
		}
	}

	//Stats.RemainingStacks = stackPool.Remainder()
	//Stats.RemainingStackPtrs = stackPtrPool.Remainder()

	//If the number of threads is greater than one, a final pass is required
	if numThreads > 1 {
		startRecombiningStacks := time.Now()
		//Create the final input by joining together the stacks from the previous step
		finalPassInput := newLos(stackPoolFinalPass)
		for i := 0; i < numThreads; i++ {
			iterator := parseResults[i].stack.HeadIterator()
			//Ignore the first token
			iterator.Next()
			sym := iterator.Next()
			for sym != nil {
				finalPassInput.Push(sym)
				sym = iterator.Next()
			}
		}
		Stats.RecombiningStacksTime = time.Since(startRecombiningStacks)

		//fmt.Print("Final pass thread input: ")
		//finalPassThreadContext.input.Println()
		//fmt.Print("Final pass thread stack: ")
		//finalPassThreadContext.stack.Println()

		go threadJob(1, 0, true, &finalPassInput, nil, stackPoolNewNonterminalsFinalPass, stackPtrPoolFinalPass, c)

		finalPassParseResult := <-c

		//fmt.Println("Final thread finished parsing")
		//fmt.Println("Result:", finalPassThreadContext.result)
		//fmt.Print("Final stack: ")
		//finalPassThreadContext.stack.Println()

		if !finalPassParseResult.success {
			Stats.ParseTimeTotal = time.Since(start)
			return nil, errors.New("Parsing error")
		}

		//Pop tokens from the stack until a nonterminal is found
		sym := finalPassParseResult.stack.Pop()

		for isTerminal(sym.Token) {
			sym = finalPassParseResult.stack.Pop()
		}

		//sym.PrintTreeln()

		//Set the result as the nonterminal symbol
		result = sym

		Stats.RemainingStacksFinalPass = stackPoolFinalPass.Remainder()
		Stats.RemainingStacksNewNonterminalsFinalPass = stackPoolNewNonterminalsFinalPass.Remainder()
		Stats.RemainingStackPtrsFinalPass = stackPtrPoolFinalPass.Remainder()
	} else {
		//Pop tokens from the stack until a nonterminal is found
		sym := parseResults[0].stack.Pop()

		for isTerminal(sym.Token) {
			sym = parseResults[0].stack.Pop()
		}

		//sym.PrintTreeln()

		//Set the result as the nonterminal symbol
		result = sym
	}

	Stats.RemainingStacks = make([]int, numThreads)
	Stats.RemainingStacksNewNonterminals = make([]int, numThreads)
	Stats.RemainingStackPtrs = make([]int, numThreads)

	for i := 0; i < numThreads; i++ {
		Stats.RemainingStacks[i] = stackPools[i].Remainder()
		Stats.RemainingStacksNewNonterminals[i] = stackPoolsNewNonterminals[i].Remainder()
		Stats.RemainingStackPtrs[i] = stackPtrPools[i].Remainder()
	}

	Stats.ParseTimeTotal = time.Since(start)

	//Stats.RemainingStacks = stackPool.Remainder()
	//Stats.RemainingStackPtrs = stackPtrPool.Remainder()

	return result, nil
}

/*
ParseFile parses a file in parallel using an operator precedence grammar.
It takes as input a filename and the number of threads, and returns a boolean
representing the success or failure of the parsing and the symbol at the root of the syntactic tree (if successful).
*/
func ParseFile(filename string, numThreads int) (*symbol, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return ParseString(bytes, numThreads)
}
