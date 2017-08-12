%cut <

LBRACKET <
RBRACKET >
LSLASH   </
RSLASH   />
EQUALS   =
VALUE    ".*"
NUMBER   [0-9]+
IDENT    [a-zA-Z0-9_\-\."=][a-zA-Z0-9_\-\."= ~]*
INFOS    [a-zA-Z0-9_:#/,&;\.=\-\(\)@?\[\]\+'$"!][a-zA-Z0-9_:#/,&;\.=\-\(\)@?\[\]\+'$"! ~]*
SPACE    [ \t]+
NEWLINE  [\r\n]

%%

{INFOS}
{
	*genSym = symbol{infos, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{LBRACKET}{IDENT}{RBRACKET}
{
	*genSym = symbol{openbracket, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{LSLASH}{IDENT}{RBRACKET}
{
	*genSym = symbol{closebracket, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{LBRACKET}{IDENT}{RSLASH}
{
	*genSym = symbol{alternativeclose, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{LBRACKET}{IDENT}{SPACE}{IDENT}{EQUALS}{VALUE}{RBRACKET}
{
	*genSym = symbol{openparams, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{LBRACKET}{IDENT}{RBRACKET}{LSLASH}{IDENT}{RBRACKET}
{
	*genSym = symbol{opencloseinfo, 0, nil, nil, nil}
	return _LEX_CORRECT
}
{LBRACKET}{IDENT}{SPACE}{IDENT}{EQUALS}{VALUE}{RBRACKET}{LSLASH}{IDENT}{RBRACKET}
{
	*genSym = symbol{opencloseparam, 0, nil, nil, nil}
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
<?[^?]+?>
{
	return _SKIP
}
.
{
	return _ERROR
}

%%
/*
lexerPreallocMem initializes all the memory pools required by the lexer.
*/
func lexerPreallocMem(inputSize int, numThreads int) {
}