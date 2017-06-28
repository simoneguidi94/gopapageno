package arithmetic

const _NUM_NONTERMINALS = 4
const _NUM_TERMINALS = 6

const (
	_S      = iota
	_E      = iota
	_T      = iota
	_F      = iota
	_PLUS   = 0x8000 + iota - _NUM_NONTERMINALS
	_TIMES  = 0x8000 + iota - _NUM_NONTERMINALS
	_LPAR   = 0x8000 + iota - _NUM_NONTERMINALS
	_RPAR   = 0x8000 + iota - _NUM_NONTERMINALS
	_NUMBER = 0x8000 + iota - _NUM_NONTERMINALS
	_TERM   = 0x8000 + iota - _NUM_NONTERMINALS
)

func tokenValue(token uint16) uint16 {
	return 0x7FFF & token
}

func isTerminal(token uint16) bool {
	return token >= 0x800
}

func tokenToString(token uint16) string {
	switch token {
	case _S:
		return "S"
	case _E:
		return "E"
	case _T:
		return "T"
	case _F:
		return "F"
	case _PLUS:
		return "PLUS"
	case _TIMES:
		return "TIMES"
	case _LPAR:
		return "LPAR"
	case _RPAR:
		return "RPAR"
	case _NUMBER:
		return "NUMBER"
	case _TERM:
		return "#"
	}
	return "UNKNOWN_TOKEN"
}
