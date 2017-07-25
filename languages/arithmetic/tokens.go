package arithmetic

const _NUM_NONTERMINALS = 5
const _NUM_TERMINALS = 6

const (
	E_F_S_T = iota
	E_S = iota
	E_S_T = iota
	NEW_AXIOM = iota
	_EMPTY = iota
	LPAR = 0x8000 + iota - _NUM_NONTERMINALS
	NUMBER = 0x8000 + iota - _NUM_NONTERMINALS
	PLUS = 0x8000 + iota - _NUM_NONTERMINALS
	RPAR = 0x8000 + iota - _NUM_NONTERMINALS
	TIMES = 0x8000 + iota - _NUM_NONTERMINALS
	_TERM = 0x8000 + iota - _NUM_NONTERMINALS
)

func tokenValue(token uint16) uint16 {
	return 0x7FFF & token
}

func isTerminal(token uint16) bool {
	return token >= 0x800
}

func tokenToString(token uint16) string {
	switch token {
	case E_F_S_T:
		return "E_F_S_T"
	case E_S:
		return "E_S"
	case E_S_T:
		return "E_S_T"
	case NEW_AXIOM:
		return "NEW_AXIOM"
	case _EMPTY:
		return "_EMPTY"
	case LPAR:
		return "LPAR"
	case NUMBER:
		return "NUMBER"
	case PLUS:
		return "PLUS"
	case RPAR:
		return "RPAR"
	case TIMES:
		return "TIMES"
	case _TERM:
		return "_TERM"
	}
	return "UNKNOWN_TOKEN"
}