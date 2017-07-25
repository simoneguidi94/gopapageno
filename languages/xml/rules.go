package xml

/*
rule represents a grammar rule of the language. Its lhs is a single token while its rhs is a slice of tokens.
*/
type rule struct {
	lhs uint16
	rhs []uint16
}

/*
The maximum length of the rhs of a rule of the language
*/
const _MAX_RHS_LEN = 4

/*
The rules of the language. They are sorted by their rhs
*/
var _RULES = []rule{
	rule{NEW_AXIOM, []uint16{ELEM}},
	rule{ELEM, []uint16{ELEM, alternativeclose}},
	rule{ELEM, []uint16{ELEM, openbracket, ELEM, closebracket}},
	rule{ELEM, []uint16{ELEM, opencloseinfo}},
	rule{ELEM, []uint16{ELEM, opencloseparam}},
	rule{ELEM, []uint16{ELEM, openparams, ELEM, closeparams}},
	rule{ELEM, []uint16{alternativeclose}},
	rule{ELEM, []uint16{infos}},
	rule{ELEM, []uint16{openbracket, ELEM, closebracket}},
	rule{ELEM, []uint16{opencloseinfo}},
	rule{ELEM, []uint16{opencloseparam}},
	rule{ELEM, []uint16{openparams, ELEM, closebracket}},
}

var compressedTrie = []uint16{2, 0, 7, 0, 17, 32769, 65, 32772, 68, 32773, 71, 32774, 84, 32775, 87, 32776, 90, 1, 0, 5, 32769, 30, 32773, 33, 32774, 46, 32775, 49, 32776, 52, 0, 1, 0, 2, 0, 1, 0, 38, 2, 0, 1, 32770, 43, 0, 2, 0, 0, 3, 0, 0, 4, 0, 2, 0, 1, 0, 57, 2, 0, 1, 32771, 62, 0, 5, 0, 0, 6, 0, 0, 7, 0, 2, 0, 1, 0, 76, 2, 0, 1, 32770, 81, 0, 8, 0, 0, 9, 0, 0, 10, 0, 2, 0, 1, 0, 95, 2, 0, 1, 32770, 100, 0, 11, 0}

/*
findMatch tries to find a match for the rhs using the compressed trie above.
On success it returns the corresponding lhs and the rule number.
On failure it returns an error.
*/
func findMatch(rhs []uint16) (uint16, uint16) {
	pos := uint16(0)

	for _, key := range rhs {
		//Skip the value and rule num for each node (except the last)
		pos += 2
		numIndices := compressedTrie[pos]
		if numIndices == 0 {
			return _EMPTY, 0
		}
		pos++
		low := uint16(0)
		high := uint16(numIndices - 1)
		startPos := pos
		foundNext := false

		for low <= high {
			indexpos := low + (high-low)/2
			pos = startPos + indexpos*2
			curKey := compressedTrie[pos]

			if key < curKey {
				high = indexpos - 1
			} else if key > curKey {
				low = indexpos + 1
			} else {
				pos = compressedTrie[pos+1]
				foundNext = true
				break
			}
		}
		if !foundNext {
			return _EMPTY, 0
		}
	}

	return compressedTrie[pos], compressedTrie[pos+1]
}