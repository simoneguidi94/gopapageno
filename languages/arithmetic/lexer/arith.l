%cut \n

LPAR    \(
RPAR    \)
PLUS    \+
TIMES   \*
DIGIT   [0-9]
SPACE   [ \t]
NEWLINE [\r\n]

%%

{LPAR} 
{
	*genSym = symbol{LPAR, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{RPAR}
{
	*genSym = symbol{RPAR, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{TIMES}
{
	*genSym = symbol{TIMES, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{PLUS}
{
	*genSym = symbol{PLUS, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{DIGIT}+
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
{SPACE}
{
	return _SKIP
}
{NEWLINE}
{
	return _SKIP
}
.
{
	return _ERROR
}

%%
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