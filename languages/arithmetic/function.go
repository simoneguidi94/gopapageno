package arithmetic

import (
	"math"
)

var parserInt64Pools []*int64Pool

/*
parserPreallocMem initializes all the memory pools required by the semantic function of the parser.
*/
func parserPreallocMem(inputSize int, numThreads int) {
	parserInt64Pools = make([]*int64Pool, numThreads)
	
	avgCharsPerNumber := float64(4)
	
	poolSizePerThread := int(math.Ceil((float64(inputSize) / avgCharsPerNumber) / float64(numThreads)))

	for i := 0; i < numThreads; i++ {
		parserInt64Pools[i] = newInt64Pool(poolSizePerThread)
	}
}

/*
function is the semantic function of the parser.
*/
func function(thread int, ruleNum uint16, lhs *symbol, rhs []*symbol) {
	switch ruleNum {
	case 0:
		NEW_AXIOM0 := lhs
		E_F_S_T1 := rhs[0]

		NEW_AXIOM0.Child = E_F_S_T1

		{
			NEW_AXIOM0.Value = E_F_S_T1.Value
		}
	case 1:
		E_S0 := lhs
		E_F_S_T1 := rhs[0]
		PLUS2 := rhs[1]
		E_F_S_T3 := rhs[2]

		E_S0.Child = E_F_S_T1
		E_F_S_T1.Next = PLUS2
		PLUS2.Next = E_F_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_F_S_T1.Value.(*int64) + *E_F_S_T3.Value.(*int64)
			E_S0.Value = newValue
		}
	case 2:
		E_S0 := lhs
		E_F_S_T1 := rhs[0]
		PLUS2 := rhs[1]
		E_S_T3 := rhs[2]

		E_S0.Child = E_F_S_T1
		E_F_S_T1.Next = PLUS2
		PLUS2.Next = E_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_F_S_T1.Value.(*int64) + *E_S_T3.Value.(*int64)
			E_S0.Value = newValue
		}
	case 3:
		E_S_T0 := lhs
		E_F_S_T1 := rhs[0]
		TIMES2 := rhs[1]
		E_F_S_T3 := rhs[2]

		E_S_T0.Child = E_F_S_T1
		E_F_S_T1.Next = TIMES2
		TIMES2.Next = E_F_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_F_S_T1.Value.(*int64) * *E_F_S_T3.Value.(*int64)
			E_S_T0.Value = newValue
		}
	case 4:
		NEW_AXIOM0 := lhs
		E_S1 := rhs[0]

		NEW_AXIOM0.Child = E_S1

		{
			NEW_AXIOM0.Value = E_S1.Value
		}
	case 5:
		E_S0 := lhs
		E_S1 := rhs[0]
		PLUS2 := rhs[1]
		E_F_S_T3 := rhs[2]

		E_S0.Child = E_S1
		E_S1.Next = PLUS2
		PLUS2.Next = E_F_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_S1.Value.(*int64) + *E_F_S_T3.Value.(*int64)
			E_S0.Value = newValue
		}
	case 6:
		E_S0 := lhs
		E_S1 := rhs[0]
		PLUS2 := rhs[1]
		E_S_T3 := rhs[2]

		E_S0.Child = E_S1
		E_S1.Next = PLUS2
		PLUS2.Next = E_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_S1.Value.(*int64) + *E_S_T3.Value.(*int64)
			E_S0.Value = newValue
		}
	case 7:
		NEW_AXIOM0 := lhs
		E_S_T1 := rhs[0]

		NEW_AXIOM0.Child = E_S_T1

		{
			NEW_AXIOM0.Value = E_S_T1.Value
		}
	case 8:
		E_S0 := lhs
		E_S_T1 := rhs[0]
		PLUS2 := rhs[1]
		E_F_S_T3 := rhs[2]

		E_S0.Child = E_S_T1
		E_S_T1.Next = PLUS2
		PLUS2.Next = E_F_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_S_T1.Value.(*int64) + *E_F_S_T3.Value.(*int64)
			E_S0.Value = newValue
		}
	case 9:
		E_S0 := lhs
		E_S_T1 := rhs[0]
		PLUS2 := rhs[1]
		E_S_T3 := rhs[2]

		E_S0.Child = E_S_T1
		E_S_T1.Next = PLUS2
		PLUS2.Next = E_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_S_T1.Value.(*int64) + *E_S_T3.Value.(*int64)
			E_S0.Value = newValue
		}
	case 10:
		E_S_T0 := lhs
		E_S_T1 := rhs[0]
		TIMES2 := rhs[1]
		E_F_S_T3 := rhs[2]

		E_S_T0.Child = E_S_T1
		E_S_T1.Next = TIMES2
		TIMES2.Next = E_F_S_T3

		{
			newValue := parserInt64Pools[thread].Get()
			*newValue = *E_S_T1.Value.(*int64) * *E_F_S_T3.Value.(*int64)
			E_S_T0.Value = newValue
		}
	case 11:
		E_F_S_T0 := lhs
		LPAR1 := rhs[0]
		E_F_S_T2 := rhs[1]
		RPAR3 := rhs[2]

		E_F_S_T0.Child = LPAR1
		LPAR1.Next = E_F_S_T2
		E_F_S_T2.Next = RPAR3

		{
			E_F_S_T0.Value = E_F_S_T2.Value
		}
	case 12:
		E_F_S_T0 := lhs
		LPAR1 := rhs[0]
		E_S2 := rhs[1]
		RPAR3 := rhs[2]

		E_F_S_T0.Child = LPAR1
		LPAR1.Next = E_S2
		E_S2.Next = RPAR3

		{
			E_F_S_T0.Value = E_S2.Value
		}
	case 13:
		E_F_S_T0 := lhs
		LPAR1 := rhs[0]
		E_S_T2 := rhs[1]
		RPAR3 := rhs[2]

		E_F_S_T0.Child = LPAR1
		LPAR1.Next = E_S_T2
		E_S_T2.Next = RPAR3

		{
			E_F_S_T0.Value = E_S_T2.Value
		}
	case 14:
		E_F_S_T0 := lhs
		NUMBER1 := rhs[0]

		E_F_S_T0.Child = NUMBER1

		{
			E_F_S_T0.Value = NUMBER1.Value
		}
	}
}
