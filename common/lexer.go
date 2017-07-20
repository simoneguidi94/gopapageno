import (
	"errors"
	"unsafe"
)

const (
	_ERROR       = -1
	_END_OF_FILE = 0
	_LEX_CORRECT = 1
	_SKIP        = 2
)

/*
lexer contains the file data and the current position.
*/
type lexer struct {
	data []byte
	pos  int
}

/*
yyLex reads a token from the lexer and creates a new symbol, saving it in genSym.
It returns one of the following codes:
ERROR       if a token couldn't be read
END_OF_FILE if there's no more data to lex
LEX_CORRECT if a token was successfully read
*/
func (l *lexer) yyLex(genSym *symbol) int {
	result := _SKIP
	for result == _SKIP {
		var lastFinalStateReached *lexerDfaState = nil
		var lastFinalStatePos int
		startPos := l.pos
		curState := &lexerAutomaton[0]
		for true {
			if l.pos == len(l.data) {
				return _END_OF_FILE
			}

			curStateIndex := curState.Transitions[l.data[l.pos]]

			//Cannot read chars anymore
			if curStateIndex == -1 {
				if lastFinalStateReached == nil {
					return _ERROR
				} else {
					l.pos = lastFinalStatePos + 1
					ruleNum := lastFinalStateReached.AssociatedRules[0]
					textBytes := l.data[startPos:l.pos]
					//TODO should be changed to safe code when Go supports no-op []byte to string conversion
					text := *(*string)(unsafe.Pointer(&textBytes))
					//fmt.Printf("%s: %d\n", text, ruleNum)
					result = lexerFunction(0, ruleNum, text, genSym)
					break
				}
			} else {
				curState = &lexerAutomaton[curStateIndex]
				if curState.IsFinal {
					lastFinalStateReached = curState
					lastFinalStatePos = l.pos
					if l.pos == len(l.data)-1 {
						l.pos = lastFinalStatePos + 1
						ruleNum := lastFinalStateReached.AssociatedRules[0]
						textBytes := l.data[startPos:l.pos]
						//TODO should be changed to safe code when Go supports no-op []byte to string conversion
						text := *(*string)(unsafe.Pointer(&textBytes))
						//fmt.Printf("%s: %d\n", text, ruleNum)
						result = lexerFunction(0, ruleNum, text, genSym)
						break
					}
				}
			}
			l.pos++
		}
	}
	return result
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