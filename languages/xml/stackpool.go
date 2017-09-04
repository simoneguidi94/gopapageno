package xml

import (
	"fmt"
)

/*
stackPool allows to preallocate elements with type stack.
It is thread-safe.
*/
type stackPool struct {
	pool []stack
	cur  int
}

/*
newStackPool creates a new pool, allocating memory for a number of elements equal to length.
*/
func newStackPool(length int) *stackPool {
	p := stackPool{make([]stack, length), 0}

	return &p
}

/*
Get gets an item from the pool if available, otherwise it initializes a new one.
It is NOT thread-safe.
*/
func (p *stackPool) Get() *stack {
	if p.cur >= len(p.pool) {
		fmt.Println("Allocating a new stack!")
		return new(stack)
	}
	addr := &p.pool[p.cur]
	p.cur++
	return addr
}

/*
Remainder returns the number of items remaining in the pool.
*/
func (p *stackPool) Remainder() int {
	return len(p.pool) - p.cur
}
