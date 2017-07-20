/*
parserPreallocMem initializes all the memory pools required by the semantic function of the parser.
*/
func parserPreallocMem(inputSize int, numThreads int) {
}
%%

%axiom ELEM

%%

ELEM : ELEM openbracket ELEM closebracket
{
} | ELEM openparams ELEM closeparams
{
} | ELEM opencloseinfo
{
} | ELEM opencloseparam
{
} | ELEM alternativeclose
{
} | openbracket ELEM closebracket
{
} | openparams ELEM closebracket
{
} | opencloseinfo
{
} | opencloseparam
{
} | alternativeclose
{
} | infos
{
};