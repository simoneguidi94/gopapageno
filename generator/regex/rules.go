package regex

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
const _MAX_RHS_LEN = 6

/*
The rules of the language. They are sorted by their rhs
*/
var _RULES = []rule{
	rule{NEW_AXIOM, []uint16{CONCATENATION_RE_SIMPLE_RE}},
	rule{RE_UNION, []uint16{CONCATENATION_RE_SIMPLE_RE, pipe, CONCATENATION_RE_SIMPLE_RE}},
	rule{RE_UNION, []uint16{CONCATENATION_RE_SIMPLE_RE, pipe, RE_SIMPLE_RE}},
	rule{NEW_AXIOM, []uint16{RE_SIMPLE_RE}},
	rule{RE_UNION, []uint16{RE_SIMPLE_RE, pipe, CONCATENATION_RE_SIMPLE_RE}},
	rule{RE_UNION, []uint16{RE_SIMPLE_RE, pipe, RE_SIMPLE_RE}},
	rule{NEW_AXIOM, []uint16{RE_UNION}},
	rule{RE_UNION, []uint16{RE_UNION, pipe, CONCATENATION_RE_SIMPLE_RE}},
	rule{RE_UNION, []uint16{RE_UNION, pipe, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{any}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{any, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{any, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{any, plus}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{any, plus, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{any, plus, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{any, star}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{any, star, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{any, star, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{char}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{char, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{char, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{char, plus}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{char, plus, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{char, plus, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{char, star}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{char, star, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{char, star, RE_SIMPLE_RE}},
	rule{SET_ITEMS, []uint16{charinset}},
	rule{SET_ITEMS, []uint16{charinset, SET_ITEMS}},
	rule{SET_ITEMS, []uint16{charinset, dash, charinset}},
	rule{SET_ITEMS, []uint16{charinset, dash, charinset, SET_ITEMS}},
	rule{RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, plus}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, plus, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, plus, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, star}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, star, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, CONCATENATION_RE_SIMPLE_RE, rpar, star, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, plus}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, plus, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, plus, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, star}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, star, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_SIMPLE_RE, rpar, star, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, plus}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, plus, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, plus, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, star}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, star, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{lpar, RE_UNION, rpar, star, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, plus}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, plus, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, plus, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, star}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, star, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, SET_ITEMS, squarerpar, star, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, plus}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, plus, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, plus, RE_SIMPLE_RE}},
	rule{RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, star}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, star, CONCATENATION_RE_SIMPLE_RE}},
	rule{CONCATENATION_RE_SIMPLE_RE, []uint16{squarelpar, caret, SET_ITEMS, squarerpar, star, RE_SIMPLE_RE}},
}

var compressedTrie = []uint16{5, 0, 8, 0, 19, 2, 37, 3, 55, 32769, 73, 32771, 116, 32772, 159, 32774, 182, 32778, 335, 1, 0, 1, 32775, 24, 5, 0, 2, 0, 31, 2, 34, 3, 1, 0, 3, 2, 0, 1, 3, 1, 32775, 42, 5, 0, 2, 0, 49, 2, 52, 3, 4, 0, 3, 5, 0, 1, 6, 1, 32775, 60, 5, 0, 2, 0, 67, 2, 70, 3, 7, 0, 3, 8, 0, 2, 9, 4, 0, 84, 2, 87, 32776, 90, 32780, 103, 0, 10, 0, 0, 11, 0, 2, 12, 2, 0, 97, 2, 100, 0, 13, 0, 0, 14, 0, 2, 15, 2, 0, 110, 2, 113, 0, 16, 0, 0, 17, 0, 2, 18, 4, 0, 127, 2, 130, 32776, 133, 32780, 146, 0, 19, 0, 0, 20, 0, 2, 21, 2, 0, 140, 2, 143, 0, 22, 0, 0, 23, 0, 2, 24, 2, 0, 153, 2, 156, 0, 25, 0, 0, 26, 0, 4, 27, 2, 4, 166, 32773, 169, 4, 28, 0, 5, 0, 1, 32772, 174, 4, 29, 1, 4, 179, 4, 30, 0, 5, 0, 3, 0, 191, 2, 239, 3, 287, 5, 0, 1, 32777, 196, 2, 31, 4, 0, 207, 2, 210, 32776, 213, 32780, 226, 0, 32, 0, 0, 33, 0, 2, 34, 2, 0, 220, 2, 223, 0, 35, 0, 0, 36, 0, 2, 37, 2, 0, 233, 2, 236, 0, 38, 0, 0, 39, 0, 5, 0, 1, 32777, 244, 2, 40, 4, 0, 255, 2, 258, 32776, 261, 32780, 274, 0, 41, 0, 0, 42, 0, 2, 43, 2, 0, 268, 2, 271, 0, 44, 0, 0, 45, 0, 2, 46, 2, 0, 281, 2, 284, 0, 47, 0, 0, 48, 0, 5, 0, 1, 32777, 292, 2, 49, 4, 0, 303, 2, 306, 32776, 309, 32780, 322, 0, 50, 0, 0, 51, 0, 2, 52, 2, 0, 316, 2, 319, 0, 53, 0, 0, 54, 0, 2, 55, 2, 0, 329, 2, 332, 0, 56, 0, 0, 57, 0, 5, 0, 2, 4, 342, 32770, 390, 5, 0, 1, 32779, 347, 2, 58, 4, 0, 358, 2, 361, 32776, 364, 32780, 377, 0, 59, 0, 0, 60, 0, 2, 61, 2, 0, 371, 2, 374, 0, 62, 0, 0, 63, 0, 2, 64, 2, 0, 384, 2, 387, 0, 65, 0, 0, 66, 0, 5, 0, 1, 4, 395, 5, 0, 1, 32779, 400, 2, 67, 4, 0, 411, 2, 414, 32776, 417, 32780, 430, 0, 68, 0, 0, 69, 0, 2, 70, 2, 0, 424, 2, 427, 0, 71, 0, 0, 72, 0, 2, 73, 2, 0, 437, 2, 440, 0, 74, 0, 0, 75, 0}

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