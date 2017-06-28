package arithmetic

import (
	"errors"
	"io/ioutil"
	"strconv"
)

const (
	_ERROR       = -1
	_END_OF_FILE = 0
	_LEX_CORRECT = 1
)

/*
lexer contains the file data and the current position.
*/
type lexer struct {
	data []byte
	pos  int
}

const _INT64_LEXER_PREALLOC_LEN int = 1000000

var integerPool *int64Pool

/*
lexerPreallocMem initializes all the memory pools required by the lexer.
*/
func lexerPreallocMem(numThreads int) {
	integerPool = newInt64Pool(_INT64_LEXER_PREALLOC_LEN)
	/*int64Pools = make([]*Int64Pool, numThreads)

	for i := 0; i < numThreads; i++ {
		int64Pools[i] = NewInt64Pool(INT64_PREALLOC_LEN)
	}*/
}

/*
yyLex reads a token from the lexer and creates a new token, saving it in genSym.
It returns one of the following codes:
ERROR       if a token couldn't be read
END_OF_FILE if there's no more data to lex
LEX_CORRECT if a token was successfully read
*/
func (l *lexer) yyLex(genSym *symbol) int {
	for l.pos < len(l.data) {
		curChar := l.data[l.pos]

		l.pos++

		switch {
		case curChar == '(':
			*genSym = symbol{_LPAR, 0 /*'('*/, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == ')':
			*genSym = symbol{_RPAR, 0 /*')'*/, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '*':
			*genSym = symbol{_TIMES, 0 /*'*'*/, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '+':
			*genSym = symbol{_PLUS, 0 /*'+'*/, nil, nil, nil}
			return _LEX_CORRECT
		case curChar >= '0' && curChar <= '9':
			startPos := l.pos - 1
			for l.pos < len(l.data) && l.data[l.pos] >= '0' && l.data[l.pos] <= '9' {
				l.pos++
			}
			num := integerPool.Get()
			err := error(nil)
			*num, err = strconv.ParseInt(string(l.data[startPos:l.pos]), 10, 64)
			if err != nil {
				return _ERROR
			}
			*genSym = symbol{_NUMBER, 0, num, nil, nil}
			return _LEX_CORRECT
		case curChar == ' ', curChar == '\t', curChar == '\n', curChar == '\r':
			//Do nothing
		default:
			return _ERROR
		}
	}
	return _END_OF_FILE
}

/*
lex reads the content of a file and lexes it, pushing each symbol in a listOfStacks.
It returns a listOfStacks containing all the lexed symbols.
An error is returned if the file couldn't be opened or it contains invalid data.
*/
func lex(filename string, stackPool *stackPool) (listOfStacks, error) {
	los := newLos(stackPool)

	//Load the whole file in memory
	fileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		return los, err
	}

	sym := symbol{}

	lexer := lexer{fileContent, 0}

	//Lex the first symbol
	res := lexer.yyLex(&sym)

	//Keep lexing until the end of the file is reached or an error occurs
	for res != _END_OF_FILE {
		if res == _ERROR {
			return los, errors.New("Errore nel lexing")
		}
		los.Push(&sym)

		res = lexer.yyLex(&sym)
	}

	return los, nil
}
