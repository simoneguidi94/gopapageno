package arithmetic

const INT64_PARSER_PREALLOC_LEN int = 1000000

var int64Pools []*int64Pool

/*
parserPreallocMem initializes all the memory pools required by the semantic function.
*/
func parserPreallocMem(numThreads int) {
	int64Pools = make([]*int64Pool, numThreads)

	for i := 0; i < numThreads; i++ {
		int64Pools[i] = newInt64Pool(INT64_PARSER_PREALLOC_LEN)
	}
}

/*
function is the semantic function of the parser.
*/
func function(thread int, ruleNum int, lhs *symbol, rhs []*symbol) {
	switch ruleNum {
	case 0:
		S0 := lhs
		E1 := rhs[0]

		S0.Child = E1

		S0.Value = E1.Value
	case 1:
		E0 := lhs
		E1 := rhs[0]
		PLUS2 := rhs[1]
		T3 := rhs[2]

		E0.Child = E1
		E1.Next = PLUS2
		PLUS2.Next = T3

		newValue := int64Pools[thread].Get()
		*newValue = *E1.Value.(*int64) + *T3.Value.(*int64)
		E0.Value = newValue
	case 2:
		E0 := lhs
		E1 := rhs[0]
		PLUS2 := rhs[1]
		F3 := rhs[2]

		E0.Child = E1
		E1.Next = PLUS2
		PLUS2.Next = F3

		newValue := int64Pools[thread].Get()
		*newValue = *E1.Value.(*int64) + *F3.Value.(*int64)
		E0.Value = newValue
	case 3:
		E0 := lhs
		T1 := rhs[0]
		PLUS2 := rhs[1]
		T3 := rhs[2]

		E0.Child = T1
		T1.Next = PLUS2
		PLUS2.Next = T3

		newValue := int64Pools[thread].Get()
		*newValue = *T1.Value.(*int64) + *T3.Value.(*int64)
		E0.Value = newValue
	case 4:
		E0 := lhs
		T1 := rhs[0]
		PLUS2 := rhs[1]
		F3 := rhs[2]

		E0.Child = T1
		T1.Next = PLUS2
		PLUS2.Next = F3

		newValue := int64Pools[thread].Get()
		*newValue = *T1.Value.(*int64) + *F3.Value.(*int64)
		E0.Value = newValue
	case 5:
		T0 := lhs
		T1 := rhs[0]
		TIMES2 := rhs[1]
		F3 := rhs[2]

		T0.Child = T1
		T1.Next = TIMES2
		TIMES2.Next = F3

		newValue := int64Pools[thread].Get()
		*newValue = *T1.Value.(*int64) * *F3.Value.(*int64)
		T0.Value = newValue
	case 6:
		E0 := lhs
		F1 := rhs[0]
		PLUS2 := rhs[1]
		T3 := rhs[2]

		E0.Child = F1
		F1.Next = PLUS2
		PLUS2.Next = T3

		newValue := int64Pools[thread].Get()
		*newValue = *F1.Value.(*int64) + *T3.Value.(*int64)
		E0.Value = newValue
	case 7:
		E0 := lhs
		F1 := rhs[0]
		PLUS2 := rhs[1]
		F3 := rhs[2]

		E0.Child = F1
		F1.Next = PLUS2
		PLUS2.Next = F3

		newValue := int64Pools[thread].Get()
		*newValue = *F1.Value.(*int64) + *F3.Value.(*int64)
		E0.Value = newValue
	case 8:
		T0 := lhs
		F1 := rhs[0]
		TIMES2 := rhs[1]
		F3 := rhs[2]

		T0.Child = F1
		F1.Next = TIMES2
		TIMES2.Next = F3

		newValue := int64Pools[thread].Get()
		*newValue = *F1.Value.(*int64) * *F3.Value.(*int64)
		T0.Value = newValue
	case 9:
		F0 := lhs
		LPAR1 := rhs[0]
		E2 := rhs[1]
		RPAR3 := rhs[2]

		F0.Child = LPAR1
		LPAR1.Next = E2
		E2.Next = RPAR3

		F0.Value = E2.Value
	case 10:
		F0 := lhs
		LPAR1 := rhs[0]
		T2 := rhs[1]
		RPAR3 := rhs[2]

		F0.Child = LPAR1
		LPAR1.Next = T2
		T2.Next = RPAR3

		F0.Value = T2.Value
	case 11:
		F0 := lhs
		LPAR1 := rhs[0]
		F2 := rhs[1]
		RPAR3 := rhs[2]

		F0.Child = LPAR1
		LPAR1.Next = F2
		F2.Next = RPAR3

		F0.Value = F2.Value
	case 12:
		F0 := lhs
		NUMBER1 := rhs[0]

		F0.Child = NUMBER1

		F0.Value = NUMBER1.Value
	}
}
