package xml

//This is approx. 1MB per stack (on 64 bit architecture)
const _STACK_PTR_SIZE int = 131000

/*
stackPtr contains a fixed size array of symbol pointers, the current position in the stack
and pointers to the previous and next stacks.
*/
type stackPtr struct {
	Data [_STACK_PTR_SIZE]*symbol
	Tos  int
	Prev *stackPtr
	Next *stackPtr
}
