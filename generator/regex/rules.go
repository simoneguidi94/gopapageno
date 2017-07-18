package regex

import (
	"errors"
	"fmt"
)

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
	rule{_NEW_AXIOM, []uint16{_CONCATENATION_RE_SIMPLE_RE}},
	rule{_RE_UNION, []uint16{_CONCATENATION_RE_SIMPLE_RE, _pipe, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_RE_UNION, []uint16{_CONCATENATION_RE_SIMPLE_RE, _pipe, _RE_SIMPLE_RE}},
	rule{_NEW_AXIOM, []uint16{_RE_SIMPLE_RE}},
	rule{_RE_UNION, []uint16{_RE_SIMPLE_RE, _pipe, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_RE_UNION, []uint16{_RE_SIMPLE_RE, _pipe, _RE_SIMPLE_RE}},
	rule{_NEW_AXIOM, []uint16{_RE_UNION}},
	rule{_RE_UNION, []uint16{_RE_UNION, _pipe, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_RE_UNION, []uint16{_RE_UNION, _pipe, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_any}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_any, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_any, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_any, _plus}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_any, _plus, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_any, _plus, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_any, _star}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_any, _star, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_any, _star, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_char}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_char, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_char, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_char, _plus}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_char, _plus, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_char, _plus, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_char, _star}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_char, _star, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_char, _star, _RE_SIMPLE_RE}},
	rule{_SET_ITEMS, []uint16{_charinset}},
	rule{_SET_ITEMS, []uint16{_charinset, _SET_ITEMS}},
	rule{_SET_ITEMS, []uint16{_charinset, _dash, _charinset}},
	rule{_SET_ITEMS, []uint16{_charinset, _dash, _charinset, _SET_ITEMS}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _plus}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _plus, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _plus, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _star}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _star, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _CONCATENATION_RE_SIMPLE_RE, _rpar, _star, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _plus}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _plus, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _plus, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _star}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _star, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_SIMPLE_RE, _rpar, _star, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _plus}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _plus, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _plus, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _star}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _star, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_lpar, _RE_UNION, _rpar, _star, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _plus}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _plus, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _plus, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _star}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _star, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _SET_ITEMS, _squarerpar, _star, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _plus}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _plus, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _plus, _RE_SIMPLE_RE}},
	rule{_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _star}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _star, _CONCATENATION_RE_SIMPLE_RE}},
	rule{_CONCATENATION_RE_SIMPLE_RE, []uint16{_squarelpar, _caret, _SET_ITEMS, _squarerpar, _star, _RE_SIMPLE_RE}},
}

/*
trieNodePtr consists in a (token) key and a pointer to a trieNode.
*/
type trieNodePtr struct {
	Key uint16
	Ptr *trieNode
}

/*
trieNode is a node of a trie. It has pointers to other trieNodes
and may have (if it's the terminal node of a rhs) a value (the corresponding lhs)
as well as the number of the rule (needed for the semantic function).
*/
type trieNode struct {
	HasValue bool
	Value    uint16
	RuleNum  int
	Branches []trieNodePtr
}

/*
Get obtains the node that is assigned to a certain key.
It finds it by binary search as the keys are sorted.
It returns null if no node is assigned to that key.
*/
func (trieNode *trieNode) Get(key uint16) *trieNode {
	branches := trieNode.Branches
	low := 0
	high := len(branches) - 1

	for low <= high {
		curPos := low + (high-low)/2
		curKey := branches[curPos].Key

		if key < curKey {
			high = curPos - 1
		} else if key > curKey {
			low = curPos + 1
		} else {
			return branches[curPos].Ptr
		}
	}

	return nil
}

/*
Find traverses a trie using the elements in rhs as keys,
and returns the last node on success or nil on failure.
*/
func (trieNode *trieNode) Find(rhs []uint16) *trieNode {
	curNode := trieNode
	for _, token := range rhs {
		nextNode := curNode.Get(token)
		if nextNode == nil {
			return nil
		}
		curNode = nextNode
	}
	return curNode
}

func (trieNode *trieNode) printR(curDepth int) {
	if trieNode.HasValue {
		fmt.Println(" ->", tokenToString(trieNode.Value))
	} else {
		fmt.Println()
	}
	for _, nodePtr := range trieNode.Branches {
		for i := 0; i < curDepth; i++ {
			fmt.Print("  ")
		}

		fmt.Print(tokenToString(nodePtr.Key))
		nodePtr.Ptr.printR(curDepth + 1)
	}
}

/*
Println prints a representation of the trie.
*/
func (trieRoot *trieNode) Println() {
	trieRoot.printR(0)
	fmt.Println()
}

/*
createTrie creates a trie from a set of rules and returns it.
The rhs of the rules must be sorted.
*/
func createTrie(rules []rule) trieNode {
	root := &trieNode{false, 0, 0, make([]trieNodePtr, 0)}

	for i, rule := range rules {
		curNode := root
		for j, token := range rule.rhs {
			nextNode := curNode.Get(token)
			if nextNode == nil {
				nextNode = &trieNode{false, 0, 0, make([]trieNodePtr, 0)}
				curNode.Branches = append(curNode.Branches, trieNodePtr{token, nextNode})
			}
			curNode = nextNode

			if j == len(rule.rhs)-1 {
				curNode.HasValue = true
				curNode.Value = rule.lhs
				curNode.RuleNum = i
			}
		}
	}

	return *root
}

/*
The trie used by the parser in the reduce process
*/
var trie trieNode

/*
FindMatch tries to find a match for the rhs using the variable trie which must be previously initialized.
On success it returns the corresponding lhs and the rule number.
On failure it returns an error.
*/
func findMatch(rhs []uint16) (uint16, int, error) {
	res := trie.Find(rhs)

	if res != nil && res.HasValue {
		return res.Value, res.RuleNum, nil
	}
	return 0, 0, errors.New("")
}