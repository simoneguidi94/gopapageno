package regex

const _NUM_NONTERMINALS = 5
const _NUM_TERMINALS = 13

const (
	_CONCATENATION_RE_SIMPLE_RE = iota
	_NEW_AXIOM = iota
	_RE_SIMPLE_RE = iota
	_RE_UNION = iota
	_SET_ITEMS = iota
	__TERM = 0x8000 + iota - _NUM_NONTERMINALS
	_any = 0x8000 + iota - _NUM_NONTERMINALS
	_caret = 0x8000 + iota - _NUM_NONTERMINALS
	_char = 0x8000 + iota - _NUM_NONTERMINALS
	_charinset = 0x8000 + iota - _NUM_NONTERMINALS
	_dash = 0x8000 + iota - _NUM_NONTERMINALS
	_lpar = 0x8000 + iota - _NUM_NONTERMINALS
	_pipe = 0x8000 + iota - _NUM_NONTERMINALS
	_plus = 0x8000 + iota - _NUM_NONTERMINALS
	_rpar = 0x8000 + iota - _NUM_NONTERMINALS
	_squarelpar = 0x8000 + iota - _NUM_NONTERMINALS
	_squarerpar = 0x8000 + iota - _NUM_NONTERMINALS
	_star = 0x8000 + iota - _NUM_NONTERMINALS
)

func tokenValue(token uint16) uint16 {
	return 0x7FFF & token
}

func isTerminal(token uint16) bool {
	return token >= 0x800
}

func tokenToString(token uint16) string {
	switch token {
	case _CONCATENATION_RE_SIMPLE_RE:
		return "CONCATENATION_RE_SIMPLE_RE"
	case _NEW_AXIOM:
		return "NEW_AXIOM"
	case _RE_SIMPLE_RE:
		return "RE_SIMPLE_RE"
	case _RE_UNION:
		return "RE_UNION"
	case _SET_ITEMS:
		return "SET_ITEMS"
	case __TERM:
		return "_TERM"
	case _any:
		return "any"
	case _caret:
		return "caret"
	case _char:
		return "char"
	case _charinset:
		return "charinset"
	case _dash:
		return "dash"
	case _lpar:
		return "lpar"
	case _pipe:
		return "pipe"
	case _plus:
		return "plus"
	case _rpar:
		return "rpar"
	case _squarelpar:
		return "squarelpar"
	case _squarerpar:
		return "squarerpar"
	case _star:
		return "star"
	}
	return "UNKNOWN_TOKEN"
}