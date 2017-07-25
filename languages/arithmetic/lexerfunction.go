package arithmetic

import (
	"math"
	"strconv"
)

var lexerInt64Pools []*int64Pool

/*
lexerPreallocMem initializes all the memory pools required by the lexer.
*/
func lexerPreallocMem(inputSize int, numThreads int) {
	lexerInt64Pools = make([]*int64Pool, numThreads)
	
	avgCharsPerNumber := float64(4)
	
	poolSizePerThread := int(math.Ceil((float64(inputSize) / avgCharsPerNumber) / float64(numThreads)))

	for i := 0; i < numThreads; i++ {
		lexerInt64Pools[i] = newInt64Pool(poolSizePerThread)
	}
}

/*
lexerFunction is the semantic function of the lexer.
*/
func lexerFunction(thread int, ruleNum int, yytext string, genSym *symbol) int {
	switch ruleNum {
	case 0:
		{
			*genSym = symbol{LPAR, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 1:
		{
			*genSym = symbol{RPAR, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 2:
		{
			*genSym = symbol{TIMES, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 3:
		{
			*genSym = symbol{PLUS, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 4:
		{
			num := lexerInt64Pools[thread].Get()
			err := error(nil)
			*num, err = strconv.ParseInt(yytext, 10, 64)
			if err != nil {
				return _ERROR
			}
			*genSym = symbol{NUMBER, 0, num, nil, nil}
			return _LEX_CORRECT
		}
	case 5:
		{
			return _SKIP
		}
	case 6:
		{
			return _SKIP
		}
	case 7:
		{
			return _ERROR
		}
	}
	return _ERROR
}
