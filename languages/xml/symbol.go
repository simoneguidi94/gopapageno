package xml

/*
symbol contains a token and its value, a precedence and pointers to build the syntactic tree.
*/
type symbol struct {
	Token      uint16
	Precedence uint16
	Value      interface{}
	Next       *symbol
	Child      *symbol
}

func (s *symbol) printTreeR(level int) {
	/*if s == nil {
		return
	}

	fmt.Print(">  ")
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Print(tokenToString(s.Token), ": ")
	if s.Token == _PLUS || s.Token == _TIMES || s.Token == _LPAR || s.Token == _RPAR {
		fmt.Println(string(s.Value.(int32)))
	} else {
		fmt.Println(*s.Value.(*int64))
	}
	s.Child.printTreeR(level + 1)
	s.Next.printTreeR(level)*/
}

/*
PrintTreeln prints the syntactic tree using the symbol as the root.
*/
func (root *symbol) PrintTreeln() {
	root.printTreeR(0)
}
