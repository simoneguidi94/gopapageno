package regex

const _NUM_NONTERMINALS = 6
const _NUM_TERMINALS = 13

const (
	CONCATENATION_RE_SIMPLE_RE = iota
	NEW_AXIOM = iota
	RE_SIMPLE_RE = iota
	RE_UNION = iota
	SET_ITEMS = iota
	_EMPTY = iota
	_TERM = 0x8000 + iota - _NUM_NONTERMINALS
	any = 0x8000 + iota - _NUM_NONTERMINALS
	caret = 0x8000 + iota - _NUM_NONTERMINALS
	char = 0x8000 + iota - _NUM_NONTERMINALS
	charinset = 0x8000 + iota - _NUM_NONTERMINALS
	dash = 0x8000 + iota - _NUM_NONTERMINALS
	lpar = 0x8000 + iota - _NUM_NONTERMINALS
	pipe = 0x8000 + iota - _NUM_NONTERMINALS
	plus = 0x8000 + iota - _NUM_NONTERMINALS
	rpar = 0x8000 + iota - _NUM_NONTERMINALS
	squarelpar = 0x8000 + iota - _NUM_NONTERMINALS
	squarerpar = 0x8000 + iota - _NUM_NONTERMINALS
	star = 0x8000 + iota - _NUM_NONTERMINALS
)

func tokenValue(token uint16) uint16 {
	return 0x7FFF & token
}

func isTerminal(token uint16) bool {
	return token >= 0x800
}

func tokenToString(token uint16) string {
	switch token {
	case CONCATENATION_RE_SIMPLE_RE:
		return "CONCATENATION_RE_SIMPLE_RE"
	case NEW_AXIOM:
		return "NEW_AXIOM"
	case RE_SIMPLE_RE:
		return "RE_SIMPLE_RE"
	case RE_UNION:
		return "RE_UNION"
	case SET_ITEMS:
		return "SET_ITEMS"
	case _EMPTY:
		return "_EMPTY"
	case _TERM:
		return "_TERM"
	case any:
		return "any"
	case caret:
		return "caret"
	case char:
		return "char"
	case charinset:
		return "charinset"
	case dash:
		return "dash"
	case lpar:
		return "lpar"
	case pipe:
		return "pipe"
	case plus:
		return "plus"
	case rpar:
		return "rpar"
	case squarelpar:
		return "squarelpar"
	case squarerpar:
		return "squarerpar"
	case star:
		return "star"
	}
	return "UNKNOWN_TOKEN"
}