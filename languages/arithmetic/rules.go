package arithmetic

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
const _MAX_RHS_LEN = 3

/*
The rules of the language. They are sorted by their rhs
*/
var _RULES = []rule{
	rule{NEW_AXIOM, []uint16{E_F_S_T}},
	rule{E_S, []uint16{E_F_S_T, PLUS, E_F_S_T}},
	rule{E_S, []uint16{E_F_S_T, PLUS, E_S_T}},
	rule{E_S_T, []uint16{E_F_S_T, TIMES, E_F_S_T}},
	rule{NEW_AXIOM, []uint16{E_S}},
	rule{E_S, []uint16{E_S, PLUS, E_F_S_T}},
	rule{E_S, []uint16{E_S, PLUS, E_S_T}},
	rule{NEW_AXIOM, []uint16{E_S_T}},
	rule{E_S, []uint16{E_S_T, PLUS, E_F_S_T}},
	rule{E_S, []uint16{E_S_T, PLUS, E_S_T}},
	rule{E_S_T, []uint16{E_S_T, TIMES, E_F_S_T}},
	rule{E_F_S_T, []uint16{LPAR, E_F_S_T, RPAR}},
	rule{E_F_S_T, []uint16{LPAR, E_S, RPAR}},
	rule{E_F_S_T, []uint16{LPAR, E_S_T, RPAR}},
	rule{E_F_S_T, []uint16{NUMBER}},
}

var compressedTrie = []uint16{4, 0, 5, 0, 13, 1, 41, 2, 59, 32768, 87, 32769, 120, 3, 0, 2, 32770, 20, 32772, 33, 4, 0, 2, 0, 27, 2, 30, 1, 1, 0, 1, 2, 0, 4, 0, 1, 0, 38, 2, 3, 0, 3, 4, 1, 32770, 46, 4, 0, 2, 0, 53, 2, 56, 1, 5, 0, 1, 6, 0, 3, 7, 2, 32770, 66, 32772, 79, 4, 0, 2, 0, 73, 2, 76, 1, 8, 0, 1, 9, 0, 4, 0, 1, 0, 84, 2, 10, 0, 4, 0, 3, 0, 96, 1, 104, 2, 112, 4, 0, 1, 32771, 101, 0, 11, 0, 4, 0, 1, 32771, 109, 0, 12, 0, 4, 0, 1, 32771, 117, 0, 13, 0, 0, 14, 0}

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