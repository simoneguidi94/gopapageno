package xml

import (
	"fmt"
	"math"
)

/*
listOfStacks is a list of stacks containing symbols.
When the current stack is full a new one is automatically obtained and linked to it.
*/
type listOfStacks struct {
	head *stack
	cur  *stack
	len  int
	pool *stackPool
}

/*
iterator allows to iterate over a listOfStacks, either forward or backward.
*/
type iterator struct {
	los *listOfStacks
	cur *stack
	pos int
}

/*
newLos creates a new listOfStacks initialized with one empty stack.
*/
func newLos(pool *stackPool) listOfStacks {
	firstStack := pool.Get()
	return listOfStacks{firstStack, firstStack, 0, pool}
}

/*
Push pushes a symbol in the listOfStacks.
It returns a pointer to the pushed symbol.
*/
func (l *listOfStacks) Push(sym *symbol) *symbol {
	curStack := l.cur

	//If the current stack is full obtain a new stack and set it as the current one
	if curStack.Tos < _STACK_SIZE {
		//Copy the symbol in the current position
		curStack.Data[curStack.Tos] = *sym

		//Save the pointer to the pushed symbol
		symPtr := &curStack.Data[curStack.Tos]

		//Increment the current position
		curStack.Tos++

		//Increment the total length of the list
		l.len++

		//Return the pointer to the pushed symbol
		return symPtr
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

	//Copy the symbol in the current position
	curStack.Data[curStack.Tos] = *sym

	//Save the pointer to the pushed symbol
	symPtr := &curStack.Data[curStack.Tos]

	//Increment the current position
	curStack.Tos++

	//Increment the total length of the list
	l.len++

	//Return the pointer to the pushed symbol
	return symPtr
}

/*
Pop pops a symbol from the stack and returns a pointer to it.
*/
func (l *listOfStacks) Pop() *symbol {
	curStack := l.cur

	//Decrement the current position
	curStack.Tos--

	//While the current stack is empty set the previous stack as the current one.
	//If there are no more stacks return nil
	if curStack.Tos >= 0 {
		//Decrement the total length of the list
		l.len--

		//Return the pointer to the symbol
		return &curStack.Data[curStack.Tos]
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

	//Return the pointer to the symbol
	return &curStack.Data[curStack.Tos]
}

/*
Merge merges a listOfStacks to another by linking their stacks.
*/
func (l *listOfStacks) Merge(l2 listOfStacks) {
	l.cur.Next = l2.head
	l2.head.Prev = l.cur
	l.cur = l2.cur
	l.len += l2.len
}

/*
Split splits a listOfStacks into a number of lists equal to numSplits,
which are returned as a slice of listOfStacks.
If there are not at least numSplits stacks in the listOfStacks it panics.
The original listOfStacks should not be used after this operation.
*/
func (l *listOfStacks) Split(numSplits int) []listOfStacks {
	if numSplits > l.NumStacks() {
		panic(fmt.Sprintln("Cannot apply", numSplits, "splits on a listOfStacks containing only", l.NumStacks(), "stacks."))
	}

	listsOfStacks := make([]listOfStacks, numSplits)
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
		listsOfStacks[curList] = listOfStacks{curStack, curStack, curStack.Tos, l.pool}

		for i := 1; i < stacksToAssign; i++ {
			curStack = curStack.Next
			listsOfStacks[curList].cur = curStack
			listsOfStacks[curList].len += curStack.Tos
		}
		nextStack := curStack.Next
		curStack.Next = nil
		curStack = nextStack

		remainder -= float64(stacksToAssign)
		totAssignedStacks += stacksToAssign

		curList++
	}

	return listsOfStacks
}

/*
Length returns the number of symbols contained in the listOfStacks
*/
func (l *listOfStacks) Length() int {
	return l.len
}

/*
NumStacks returns the number of stacks contained in the listOfStacks
*/
func (l *listOfStacks) NumStacks() int {
	i := 0

	curStack := l.head

	for curStack != nil {
		i++
		curStack = curStack.Next
	}

	return i
}

/*
Println prints the content of the listOfStacks.
*/
func (l *listOfStacks) Println() {
	iterator := l.HeadIterator()

	sym := iterator.Next()
	for sym != nil {
		fmt.Printf("(%s, %s) -> ", tokenToString(sym.Token), precToString(sym.Precedence))
		sym = iterator.Next()
	}
	fmt.Println()
}

/*
HeadIterator returns an iterator initialized to point before the first element of the list.
*/
func (l *listOfStacks) HeadIterator() iterator {
	return iterator{l, l.head, -1}
}

/*
TailIterator returns an iterator initialized to point after the last element of the list.
*/
func (l *listOfStacks) TailIterator() iterator {
	return iterator{l, l.cur, l.cur.Tos}
}

/*
Prev moves the iterator one position backward and returns a pointer to the current symbol.
It returns nil if it points before the first element of the list.
*/
func (i *iterator) Prev() *symbol {
	curStack := i.cur

	i.pos--

	if i.pos >= 0 {
		return &curStack.Data[i.pos]
	}

	i.pos = -1
	if curStack.Prev == nil {
		return nil
	}
	curStack = curStack.Prev
	i.cur = curStack
	i.pos = curStack.Tos - 1

	return &curStack.Data[i.pos]
}

/*
Cur returns a pointer to the current symbol.
It returns nil if it points before the first element or after the last element of the list.
*/
func (i *iterator) Cur() *symbol {
	curStack := i.cur

	if i.pos >= 0 && i.pos < curStack.Tos {
		return &curStack.Data[i.pos]
	}

	return nil
}

/*
Next moves the iterator one position forward and returns a pointer to the current symbol.
It returns nil if it points after the last element of the list.
*/
func (i *iterator) Next() *symbol {
	curStack := i.cur

	i.pos++

	if i.pos < curStack.Tos {
		return &curStack.Data[i.pos]
	}

	i.pos = curStack.Tos
	if curStack.Next == nil {
		return nil
	}
	curStack = curStack.Next
	i.cur = curStack
	i.pos = 0

	return &curStack.Data[i.pos]
}
