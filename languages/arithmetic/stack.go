package arithmetic

const _STACK_SIZE int = 1000

/*
stack contains a fixed size array of symbols, the current position in the stack
and pointers to the previous and next stacks.
*/
type stack struct {
	Data [_STACK_SIZE]symbol
	Tos  int
	Prev *stack
	Next *stack
}
