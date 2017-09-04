package arithmetic

import (
	"fmt"
)

/*
stackPtrPool allows to preallocate elements with type stackPtr.
It is thread-safe.
*/
type stackPtrPool struct {
	pool []stackPtr
	cur  int
}

/*
newStackPtrPool creates a new pool, allocating memory for a number of elements equal to length.
*/
func newStackPtrPool(length int) *stackPtrPool {
	p := stackPtrPool{make([]stackPtr, length), 0}

	return &p
}

/*
Get gets an item from the pool if available, otherwise it initializes a new one.
It is NOT thread-safe.
*/
func (p *stackPtrPool) Get() *stackPtr {
	if p.cur >= len(p.pool) {
		fmt.Println("Allocating a new stack!")
		return new(stackPtr)
	}
	addr := &p.pool[p.cur]
	p.cur++
	return addr
}

/*
Remainder returns the number of items remaining in the pool.
*/
func (p *stackPtrPool) Remainder() int {
	return len(p.pool) - p.cur
}
