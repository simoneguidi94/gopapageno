package arithmetic

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
The maximum length of a rule of the language
*/
const _ARITH_MAX_RHS_LEN = 3

/*
The rules of the language. They are sorted by their rhs
*/
var _ARITH_RULES = []rule{
	rule{_S, []uint16{_E}},
	rule{_E, []uint16{_E, _PLUS, _T}},
	rule{_E, []uint16{_E, _PLUS, _F}},
	rule{_E, []uint16{_T, _PLUS, _T}},
	rule{_E, []uint16{_T, _PLUS, _F}},
	rule{_T, []uint16{_T, _TIMES, _F}},
	rule{_E, []uint16{_F, _PLUS, _T}},
	rule{_E, []uint16{_F, _PLUS, _F}},
	rule{_T, []uint16{_F, _TIMES, _F}},
	rule{_F, []uint16{_LPAR, _E, _RPAR}},
	rule{_F, []uint16{_LPAR, _T, _RPAR}},
	rule{_F, []uint16{_LPAR, _F, _RPAR}},
	rule{_F, []uint16{_NUMBER}},
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
