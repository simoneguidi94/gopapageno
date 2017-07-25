package generator

import (
	"fmt"
	"strings"
)

type rule struct {
	LHS    string
	RHS    []string
	Action string
}

func (r rule) String() string {
	return fmt.Sprintf("%s -> %s", r.LHS, strings.Join(r.RHS, ", "))
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
		fmt.Println(" ->", trieNode.Value)
	} else {
		fmt.Println()
	}
	for _, nodePtr := range trieNode.Branches {
		for i := 0; i < curDepth; i++ {
			fmt.Print("  ")
		}

		fmt.Print(nodePtr.Key)
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
func createTrie(rules []rule, nonterminals stringSet, terminals stringSet) trieNode {
	root := &trieNode{false, 0, 0, make([]trieNodePtr, 0)}

	for i, rule := range rules {
		curNode := root
		for j, strToken := range rule.RHS {
			token := tokenToInt(strToken, nonterminals, terminals)
			nextNode := curNode.Get(token)
			if nextNode == nil {
				nextNode = &trieNode{false, 0, 0, make([]trieNodePtr, 0)}
				curNode.Branches = append(curNode.Branches, trieNodePtr{token, nextNode})
			}
			curNode = nextNode

			if j == len(rule.RHS)-1 {
				curNode.HasValue = true
				curNode.Value = tokenToInt(rule.LHS, nonterminals, terminals)
				curNode.RuleNum = i
			}
		}
	}

	return *root
}

func (trie *trieNode) compressR(newTrie *[]uint16, curpos *uint16, nonterminals stringSet, terminals stringSet) {
	//Append the value of this node if it has one, and the rule number
	if trie.HasValue {
		*newTrie = append(*newTrie, trie.Value)
		*newTrie = append(*newTrie, uint16(trie.RuleNum))
	} else {
		*newTrie = append(*newTrie, tokenToInt("_EMPTY", nonterminals, terminals))
		*newTrie = append(*newTrie, 0)
	}
	*curpos += 2

	//Append the number of indices of this node
	*newTrie = append(*newTrie, uint16(len(trie.Branches)))
	*curpos++

	startPos := *curpos
	for i := 0; i < len(trie.Branches); i++ {
		*newTrie = append(*newTrie, trie.Branches[i].Key)
		//The offset will be updated later
		*newTrie = append(*newTrie, 0)
		*curpos += 2
	}

	for i := 0; i < len(trie.Branches); i++ {
		//Update the offset
		(*newTrie)[startPos+1+uint16(i)*2] = *curpos
		//Call compress on the pointed node
		trie.Branches[i].Ptr.compressR(newTrie, curpos, nonterminals, terminals)
	}
}

func (trie *trieNode) Compress(nonterminals stringSet, terminals stringSet) []uint16 {
	compressedTrie := make([]uint16, 0)
	curPos := uint16(0)

	trie.compressR(&compressedTrie, &curPos, nonterminals, terminals)

	return compressedTrie
}

func tokenToInt(token string, nonterminals stringSet, terminals stringSet) uint16 {
	if nonterminals.Contains(token) {
		for i, t := range nonterminals {
			if token == t {
				return uint16(i)
			}
		}
	} else {
		for i, t := range terminals {
			if token == t {
				return uint16(0x8000 + i)
			}
		}
	}

	return 0
}
