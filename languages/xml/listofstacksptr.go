package xml

import (
	"fmt"
	"math"
)

/*
listOfStackPtrs is a list of stacks containing symbol pointers.
When the current stack is full a new one is automatically obtained and linked to it.
*/
type listOfStackPtrs struct {
	head          *stackPtr
	cur           *stackPtr
	len           int
	firstTerminal *symbol
	pool          *stackPtrPool
}

/*
iteratorPtr allows to iterate over a listOfStackPtrs, either forward or backward.
*/
type iteratorPtr struct {
	los *listOfStackPtrs
	cur *stackPtr
	pos int
}

/*
newLosPtr creates a new listOfStackPtrs initialized with one empty stack.
*/
func newLosPtr(pool *stackPtrPool) listOfStackPtrs {
	firstStack := pool.Get()
	return listOfStackPtrs{firstStack, firstStack, 0, nil, pool}
}

/*
Push pushes a symbol pointer in the listOfStackPtrs.
It returns the pointer itself.
*/
func (l *listOfStackPtrs) Push(sym *symbol) *symbol {
	curStack := l.cur

	//If the current stack is full obtain a new stack and set it as the current one
	if curStack.Tos < _STACK_PTR_SIZE {
		//Copy the symbol pointer in the current position
		curStack.Data[curStack.Tos] = sym

		//If the symbol is a terminal update the firstTerminal pointer
		if isTerminal(sym.Token) {
			l.firstTerminal = sym
		}

		//Increment the current position
		curStack.Tos++

		//Increment the total length of the list
		l.len++

		//Return the symbol pointer
		return sym
	}

	if curStack.Next != nil {
		curStack = curStack.Next
	} else {
		newStack := l.pool.Get()
		curStack.Next = newStack
		newStack.Prev = curStack
		curStack = newStack
	}
	l.cur = curStack

	//Copy the symbol pointer in the current position
	curStack.Data[curStack.Tos] = sym

	//If the symbol is a terminal update the firstTerminal pointer
	if isTerminal(sym.Token) {
		l.firstTerminal = sym
	}

	//Increment the current position
	curStack.Tos++

	//Increment the total length of the list
	l.len++

	//Return the symbol pointer
	return sym
}

/*
Pop pops a symbol pointer from the stack and returns it.
*/
func (l *listOfStackPtrs) Pop() *symbol {
	curStack := l.cur

	//Decrement the current position
	curStack.Tos--

	//While the current stack is empty set the previous stack as the current one.
	//If there are no more stacks return nil
	if curStack.Tos >= 0 {
		//Decrement the total length of the list
		l.len--

		//Return the symbol pointer
		return curStack.Data[curStack.Tos]
	}

	curStack.Tos = 0

	if curStack.Prev == nil {
		return nil
	}

	curStack = curStack.Prev
	curStack.Tos--
	l.cur = curStack

	//Decrement the total length of the list
	l.len--

	//Return the symbol pointer
	return curStack.Data[curStack.Tos]
}

/*
Merge merges a listOfStackPtrs to another by linking their stacks.
*/
func (l *listOfStackPtrs) Merge(l2 listOfStackPtrs) {
	l.cur.Next = l2.head
	l2.head.Prev = l.cur
	l.cur = l2.cur
	l.len += l2.len
	l.firstTerminal = l2.firstTerminal
}

/*
Split splits a listOfStackPtrs into a number of lists equal to numSplits,
which are returned as a slice of listOfStackPtrs.
If there are not at least numSplits stacks in the listOfStackPtrs it panics.
The original listOfStackPtrs should not be used after this operation.
*/
func (l *listOfStackPtrs) Split(numSplits int) []listOfStackPtrs {
	if numSplits > l.NumStacks() {
		panic(fmt.Sprintln("Cannot apply", numSplits, "splits on a ListOfStacks containing only", l.NumStacks(), "stacks."))
	}

	listsOfStacks := make([]listOfStackPtrs, numSplits)
	curList := 0

	numStacks := l.NumStacks()
	deltaStacks := float64(numStacks) / float64(numSplits)
	totAssignedStacks := 0
	remainder := float64(0)

	curStack := l.head

	for totAssignedStacks < numStacks {
		remainder += deltaStacks
		stacksToAssign := int(math.Floor(remainder + 0.5))

		curStack.Prev = nil
		listsOfStacks[curList] = listOfStackPtrs{curStack, curStack, curStack.Tos, nil, l.pool}

		for i := 1; i < stacksToAssign; i++ {
			curStack = curStack.Next
			listsOfStacks[curList].cur = curStack
			listsOfStacks[curList].len += curStack.Tos
		}
		nextStack := curStack.Next
		curStack.Next = nil
		curStack = nextStack

		listsOfStacks[curList].firstTerminal = listsOfStacks[curList].findFirstTerminal()

		remainder -= float64(stacksToAssign)
		totAssignedStacks += stacksToAssign

		curList++
	}

	return listsOfStacks
}

/*
Length returns the number of symbol pointers contained in the listOfStackPtrs
*/
func (l *listOfStackPtrs) Length() int {
	return l.len
}

/*
NumStacks returns the number of stacks contained in the listOfStackPtrs
*/
func (l *listOfStackPtrs) NumStacks() int {
	i := 0

	curStack := l.head

	for curStack != nil {
		i++
		curStack = curStack.Next
	}

	return i
}

/*
FirstTerminal returns a pointer to the first terminal on the stack.
*/
func (l *listOfStackPtrs) FirstTerminal() *symbol {
	return l.firstTerminal
}

/*
UpdateFirstTerminal should be used after a reduction in order to update the first terminal counter.
In fact, in order to save some time, only the Push operation automatically updates the first terminal pointer,
while the Pop operation does not.
*/
func (l *listOfStackPtrs) UpdateFirstTerminal() {
	l.firstTerminal = l.findFirstTerminal()
}

/*
This function is for internal usage only.
*/
func (l *listOfStackPtrs) findFirstTerminal() *symbol {
	curStack := l.cur

	pos := curStack.Tos - 1

	for pos < 0 {
		pos = -1
		if curStack.Prev == nil {
			return nil
		}
		curStack = curStack.Prev
		pos = curStack.Tos - 1
	}

	for !isTerminal(curStack.Data[pos].Token) {
		pos--
		for pos < 0 {
			pos = -1
			if curStack.Prev == nil {
				return nil
			}
			curStack = curStack.Prev
			pos = curStack.Tos - 1
		}
	}

	return curStack.Data[pos]
}

/*
Println prints the content of the listOfStackPtrs.
*/
func (l *listOfStackPtrs) Println() {
	iterator := l.HeadIterator()

	sym := iterator.Next()
	for sym != nil {
		fmt.Printf("(%s, %s) -> ", tokenToString(sym.Token), precToString(sym.Precedence))
		sym = iterator.Next()
	}
	fmt.Println()
}

/*
HeadIterator returns an Iterator starting at the first element of the list.
*/
func (l *listOfStackPtrs) HeadIterator() iteratorPtr {
	return iteratorPtr{l, l.head, -1}
}

/*
TailIterator returns an Iterator starting at the last element of the list.
*/
func (l *listOfStackPtrs) TailIterator() iteratorPtr {
	return iteratorPtr{l, l.cur, l.cur.Tos}
}

/*
Prev moves the iterator one position backward and returns a pointer to the current symbol.
It returns nil if it points before the first element of the list.
*/
func (i *iteratorPtr) Prev() *symbol {
	curStack := i.cur

	i.pos--

	if i.pos >= 0 {
		return curStack.Data[i.pos]
	}

	i.pos = -1
	if curStack.Prev == nil {
		return nil
	}
	curStack = curStack.Prev
	i.cur = curStack
	i.pos = curStack.Tos - 1

	return curStack.Data[i.pos]
}

/*
Cur returns a pointer to the current symbol.
It returns nil if it points before the first element or after the last element of the list.
*/
func (i *iteratorPtr) Cur() *symbol {
	curStack := i.cur

	if i.pos >= 0 && i.pos < curStack.Tos {
		return curStack.Data[i.pos]
	}

	return nil
}

/*
Next moves the iterator one position forward and returns a pointer to the current symbol.
It returns nil if it points after the last element of the list.
*/
func (i *iteratorPtr) Next() *symbol {
	curStack := i.cur

	i.pos++

	if i.pos < curStack.Tos {
		return curStack.Data[i.pos]
	}

	i.pos = curStack.Tos
	if curStack.Next == nil {
		return nil
	}
	curStack = curStack.Next
	i.cur = curStack
	i.pos = 0

	return curStack.Data[i.pos]
}
