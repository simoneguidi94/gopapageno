package arithmetic

/*
stackPtr contains a fixed size array of symbol pointers, the current position in the stack
and pointers to the previous and next stacks.
*/
type stackPtr struct {
	Data [_STACK_SIZE]*symbol
	Tos  int
	Prev *stackPtr
	Next *stackPtr
}
