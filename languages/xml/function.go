package xml

/*
parserPreallocMem initializes all the memory pools required by the semantic function of the parser.
*/
func parserPreallocMem(inputSize int, numThreads int) {
}

/*
function is the semantic function of the parser.
*/
func function(thread int, ruleNum uint16, lhs *symbol, rhs []*symbol) {
	switch ruleNum {
	case 0:
		NEW_AXIOM0 := lhs
		ELEM1 := rhs[0]

		NEW_AXIOM0.Child = ELEM1

		{
			NEW_AXIOM0.Value = ELEM1.Value
		}
	case 1:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		alternativeclose2 := rhs[1]

		ELEM0.Child = ELEM1
		ELEM1.Next = alternativeclose2

		{
		}
	case 2:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		openbracket2 := rhs[1]
		ELEM3 := rhs[2]
		closebracket4 := rhs[3]

		ELEM0.Child = ELEM1
		ELEM1.Next = openbracket2
		openbracket2.Next = ELEM3
		ELEM3.Next = closebracket4

		{
		}
	case 3:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		opencloseinfo2 := rhs[1]

		ELEM0.Child = ELEM1
		ELEM1.Next = opencloseinfo2

		{
		}
	case 4:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		opencloseparam2 := rhs[1]

		ELEM0.Child = ELEM1
		ELEM1.Next = opencloseparam2

		{
		}
	case 5:
		ELEM0 := lhs
		ELEM1 := rhs[0]
		openparams2 := rhs[1]
		ELEM3 := rhs[2]
		closeparams4 := rhs[3]

		ELEM0.Child = ELEM1
		ELEM1.Next = openparams2
		openparams2.Next = ELEM3
		ELEM3.Next = closeparams4

		{
		}
	case 6:
		ELEM0 := lhs
		alternativeclose1 := rhs[0]

		ELEM0.Child = alternativeclose1

		{
		}
	case 7:
		ELEM0 := lhs
		infos1 := rhs[0]

		ELEM0.Child = infos1

		{
		}
	case 8:
		ELEM0 := lhs
		openbracket1 := rhs[0]
		ELEM2 := rhs[1]
		closebracket3 := rhs[2]

		ELEM0.Child = openbracket1
		openbracket1.Next = ELEM2
		ELEM2.Next = closebracket3

		{
		}
	case 9:
		ELEM0 := lhs
		opencloseinfo1 := rhs[0]

		ELEM0.Child = opencloseinfo1

		{
		}
	case 10:
		ELEM0 := lhs
		opencloseparam1 := rhs[0]

		ELEM0.Child = opencloseparam1

		{
		}
	case 11:
		ELEM0 := lhs
		openparams1 := rhs[0]
		ELEM2 := rhs[1]
		closebracket3 := rhs[2]

		ELEM0.Child = openparams1
		openparams1.Next = ELEM2
		ELEM2.Next = closebracket3

		{
		}
	}
}
