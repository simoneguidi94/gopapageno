package arithmetic

import (
	"fmt"
	"sync"
)

/*
stackPtrPool allows to preallocate elements with type stackPtr.
It is thread-safe.
*/
type stackPtrPool struct {
	pool []stackPtr
	cur  int
	lock sync.Mutex
}

/*
newStackPtrPool creates a new pool, allocating memory for a number of elements equal to length.
*/
func newStackPtrPool(length int) *stackPtrPool {
	p := stackPtrPool{make([]stackPtr, length), 0, sync.Mutex{}}

	return &p
}

/*
GetSync gets an item from the pool if available, otherwise it initializes a new one.
It is thread-safe.
*/
func (p *stackPtrPool) GetSync() *stackPtr {
	p.lock.Lock()
	if p.cur >= len(p.pool) {
		fmt.Println("Allocating a new stack!")
		newStack := new(stackPtr)
		p.lock.Unlock()
		return newStack
	}
	addr := &p.pool[p.cur]
	p.cur++
	p.lock.Unlock()
	return addr
}
