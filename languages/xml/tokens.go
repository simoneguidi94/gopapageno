package xml

const _NUM_NONTERMINALS = 3
const _NUM_TERMINALS = 9

const (
	ELEM = iota
	NEW_AXIOM = iota
	_EMPTY = iota
	_TERM = 0x8000 + iota - _NUM_NONTERMINALS
	alternativeclose = 0x8000 + iota - _NUM_NONTERMINALS
	closebracket = 0x8000 + iota - _NUM_NONTERMINALS
	closeparams = 0x8000 + iota - _NUM_NONTERMINALS
	infos = 0x8000 + iota - _NUM_NONTERMINALS
	openbracket = 0x8000 + iota - _NUM_NONTERMINALS
	opencloseinfo = 0x8000 + iota - _NUM_NONTERMINALS
	opencloseparam = 0x8000 + iota - _NUM_NONTERMINALS
	openparams = 0x8000 + iota - _NUM_NONTERMINALS
)

func tokenValue(token uint16) uint16 {
	return 0x7FFF & token
}

func isTerminal(token uint16) bool {
	return token >= 0x800
}

func tokenToString(token uint16) string {
	switch token {
	case ELEM:
		return "ELEM"
	case NEW_AXIOM:
		return "NEW_AXIOM"
	case _EMPTY:
		return "_EMPTY"
	case _TERM:
		return "_TERM"
	case alternativeclose:
		return "alternativeclose"
	case closebracket:
		return "closebracket"
	case closeparams:
		return "closeparams"
	case infos:
		return "infos"
	case openbracket:
		return "openbracket"
	case opencloseinfo:
		return "opencloseinfo"
	case opencloseparam:
		return "opencloseparam"
	case openparams:
		return "openparams"
	}
	return "UNKNOWN_TOKEN"
}