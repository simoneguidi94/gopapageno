package xml

/*
lexerPreallocMem initializes all the memory pools required by the lexer.
*/
func lexerPreallocMem(inputSize int, numThreads int) {
}

/*
lexerFunction is the semantic function of the lexer.
*/
func lexerFunction(thread int, ruleNum int, yytext string, genSym *symbol) int {
	switch ruleNum {
	case 0:
		{
			*genSym = symbol{infos, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 1:
		{
			*genSym = symbol{openbracket, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 2:
		{
			*genSym = symbol{closebracket, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 3:
		{
			*genSym = symbol{alternativeclose, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 4:
		{
			*genSym = symbol{openparams, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 5:
		{
			*genSym = symbol{opencloseinfo, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 6:
		{
			*genSym = symbol{opencloseparam, 0, nil, nil, nil}
			return _LEX_CORRECT
		}
	case 7:
		{
			return _SKIP
		}
	case 8:
		{
			return _SKIP
		}
	case 9:
		{
			return _SKIP
		}
	case 10:
		{
			return _ERROR
		}
	}
	return _ERROR
}
