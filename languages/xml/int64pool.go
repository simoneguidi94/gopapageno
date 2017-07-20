package xml

/*
int64Pool allows to preallocate elements with type int64.
It is not thread-safe, you have to create one for each thread.
*/
type int64Pool struct {
	pool []int64
	cur  int
}

/*
newInt64Pool creates a new pool, allocating memory for a number of elements equal to length.
*/
func newInt64Pool(length int) *int64Pool {
	p := int64Pool{make([]int64, length), 0}

	for i := 0; i < length; i++ {
		p.pool[i] = int64(0)
	}

	return &p
}

/*
Get gets an item from the pool if available, otherwise it initializes a new one.
It is NOT thread-safe.
*/
func (p *int64Pool) Get() *int64 {
	if p.cur >= len(p.pool) {
		return new(int64)
	}
	addr := &p.pool[p.cur]
	p.cur++
	return addr
}

/*
Remainder returns the number of items remaining in the pool.
*/
func (p *int64Pool) Remainder() int {
	return len(p.pool) - p.cur
}