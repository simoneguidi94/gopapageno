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
%%

%axiom S

%%

S : E
{
	$$.Value = $1.Value
};

E : E PLUS T
{
	newValue := parserInt64Pools[thread].Get()
	*newValue = *$1.Value.(*int64) + *$3.Value.(*int64)
	$$.Value = newValue
} | T
{
	$$.Value = $1.Value
};

T : T TIMES F
{
	newValue := parserInt64Pools[thread].Get()
	*newValue = *$1.Value.(*int64) * *$3.Value.(*int64)
	$$.Value = newValue
} | F
{
	$$.Value = $1.Value
};

F : LPAR E RPAR
{
	$$.Value = $2.Value
} | NUMBER
{
	$$.Value = $1.Value
};
