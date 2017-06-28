package arithmetic

import (
	"fmt"
	"sync"
)

/*
stackPool allows to preallocate elements with type stack.
It is thread-safe.
*/
type stackPool struct {
	pool []stack
	cur  int
	lock sync.Mutex
}

/*
newStackPool creates a new pool, allocating memory for a number of elements equal to length.
*/
func newStackPool(length int) *stackPool {
	p := stackPool{make([]stack, length), 0, sync.Mutex{}}

	return &p
}

/*
GetSync gets an item from the pool if available, otherwise it initializes a new one.
It is thread-safe.
*/
func (p *stackPool) GetSync() *stack {
	p.lock.Lock()
	if p.cur >= len(p.pool) {
		fmt.Println("Allocating a new stack!")
		newStack := new(stack)
		p.lock.Unlock()
		return newStack
	}
	addr := &p.pool[p.cur]
	p.cur++
	p.lock.Unlock()
	return addr
}
