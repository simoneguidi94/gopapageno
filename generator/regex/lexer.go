package regex

import (
	"errors"
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

/*
lexerPreallocMem initializes all the memory pools required by the lexer.
*/
func lexerPreallocMem(inputSize int, numThreads int) {
}

var insideSet bool = false

/*
yyLex reads a token from the lexer and creates a new symbol, saving it in genSym.
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
			*genSym = symbol{lpar, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == ')':
			*genSym = symbol{rpar, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '[':
			insideSet = true
			*genSym = symbol{squarelpar, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == ']':
			insideSet = false
			*genSym = symbol{squarerpar, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '*':
			*genSym = symbol{star, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '+':
			*genSym = symbol{plus, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '-':
			*genSym = symbol{dash, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '|':
			*genSym = symbol{pipe, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '^':
			*genSym = symbol{caret, 0, nil, nil, nil}
			return _LEX_CORRECT
		case curChar == '.':
			var anyCharClass [256]bool
			//Skip the first char (empty transition)
			for i := 1; i < len(anyCharClass); i++ {
				anyCharClass[i] = true
			}
			anyCharClass['\n'] = false
			anyCharClass['\r'] = false
			newNfa := newNfaFromCharClass(anyCharClass)
			*genSym = symbol{any, 0, &newNfa, nil, nil}
			return _LEX_CORRECT
		case curChar == '\t' || curChar == '\r' || curChar == '\n':
			//Do nothing
		case curChar == '\\':
			escapedChar := l.data[l.pos]
			l.pos++
			switch escapedChar {
			case '(', ')', '[', ']', '*', '+', '-', '|', '^', '.', '\\':
				token := uint16(char)
				if insideSet {
					token = charinset
				}
				*genSym = symbol{token, 0, escapedChar, nil, nil}
				return _LEX_CORRECT
			case 't':
				token := uint16(char)
				if insideSet {
					token = charinset
				}
				*genSym = symbol{token, 0, byte('\t'), nil, nil}
				return _LEX_CORRECT
			case 'r':
				token := uint16(char)
				if insideSet {
					token = charinset
				}
				*genSym = symbol{token, 0, byte('\r'), nil, nil}
				return _LEX_CORRECT
			case 'n':
				token := uint16(char)
				if insideSet {
					token = charinset
				}
				*genSym = symbol{token, 0, byte('\n'), nil, nil}
				return _LEX_CORRECT
			default:
				return _ERROR
			}
		default:
			token := uint16(char)
			if insideSet {
				token = charinset
			}
			*genSym = symbol{token, 0, curChar, nil, nil}
			return _LEX_CORRECT
		}
	}
	return _END_OF_FILE
}

/*
lex reads an input string as a slice of byte and lexes it, pushing each symbol in a listOfStacks.
It returns a listOfStacks containing all the lexed symbols.
An error is returned if the string contains invalid data.
*/
func lex(input []byte, stackPool *stackPool) (listOfStacks, error) {
	los := newLos(stackPool)

	sym := symbol{}

	lexer := lexer{input, 0}

	//Lex the first symbol
	res := lexer.yyLex(&sym)

	//Keep lexing until the end of the file is reached or an error occurs
	for res != _END_OF_FILE {
		if res == _ERROR {
			return los, errors.New("Lexing error")
		}
		los.Push(&sym)

		res = lexer.yyLex(&sym)
	}

	return los, nil
}
