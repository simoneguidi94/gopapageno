package generator

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/simoneguidi94/gopapageno/generator/regex"
)

func emitOutputFolder(outdir string) error {
	_, err := os.Stat(outdir)

	if err != nil {
		err = os.Mkdir(outdir, os.ModeDir)
		if err != nil {
			return err
		} else {
			fmt.Println("Created directory " + outdir)
		}
	}

	return nil
}

func emitLexerAutomaton(outdir string, dfa regex.Dfa) error {
	outPath := outdir + "/" + "lexerautomaton.go"
	file, err := createFile(outPath)

	if err != nil {
		return err
	}

	defer file.Close()

	packageName := path.Base(outdir)

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	var lexerAutomaton []regex.DfaState = []regex.DfaState{
		regex.DfaState{0, [256]*regex.DfaState{}, false, nil},
	}

	_ = lexerAutomaton

	file.WriteString("var lexerAutomaton lexerDfa = []lexerDfaState {\n")
	states := dfa.GetStates()
	for _, state := range states {
		file.WriteString("\tlexerDfaState{[256]int{")
		for i := 0; i < len(state.Transitions)-1; i++ {
			if state.Transitions[i] == nil {
				file.WriteString("-1, ")
			} else {
				file.WriteString(fmt.Sprintf("%d, ", state.Transitions[i].Num))
			}
		}
		if state.Transitions[len(state.Transitions)-1] == nil {
			file.WriteString("-1")
		} else {
			file.WriteString(fmt.Sprintf("%d", state.Transitions[len(state.Transitions)-1].Num))
		}
		file.WriteString(fmt.Sprintf("}, %t, []int{", state.IsFinal))
		sort.Ints(state.AssociatedRules)
		for i := 0; i < len(state.AssociatedRules)-1; i++ {
			file.WriteString(fmt.Sprintf("%d, ", state.AssociatedRules[i]))
		}
		if len(state.AssociatedRules) > 0 {
			file.WriteString(fmt.Sprintf("%d", state.AssociatedRules[len(state.AssociatedRules)-1]))
		}
		file.WriteString("}},\n")
	}
	file.WriteString("}")

	return nil
}

func emitLexerFunction(outdir string, lexCode string, lexRules []lexRule) error {
	outPath := outdir + "/" + "lexerfunction.go"
	file, err := createFile(outPath)

	if err != nil {
		return err
	}

	defer file.Close()

	packageName := path.Base(outdir)

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	file.WriteString(lexCode)
	file.WriteString("\n\n")

	file.WriteString("/*\n")
	file.WriteString("lexerFunction is the semantic function of the lexer.\n")
	file.WriteString("*/\n")
	file.WriteString("func lexerFunction(thread int, ruleNum int, yytext string, genSym *symbol) int {\n")
	file.WriteString("\tswitch ruleNum {\n")
	for i, rule := range lexRules {
		file.WriteString(fmt.Sprintf("\tcase %d:\n", i))
		action := rule.Action
		lines := strings.Split(action, "\n")
		for _, line := range lines {
			file.WriteString("\t\t")
			file.WriteString(line)
			file.WriteString("\n")
		}
	}
	file.WriteString("\t}\n")
	file.WriteString("\treturn _ERROR\n")
	file.WriteString("}\n")

	return nil
}

func emitTokens(outdir string, nonterminals stringSet, terminals stringSet) error {
	outPath := outdir + "/" + "tokens.go"
	file, err := createFile(outPath)

	if err != nil {
		return err
	}

	defer file.Close()

	packageName := path.Base(outdir)

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	file.WriteString(fmt.Sprintf("const _NUM_NONTERMINALS = %d\n", len(nonterminals)))
	file.WriteString(fmt.Sprintf("const _NUM_TERMINALS = %d\n\n", len(terminals)))

	file.WriteString("const (\n")
	for _, token := range nonterminals {
		file.WriteString(fmt.Sprintf("\t%s = iota\n", token))
	}
	for _, token := range terminals {
		file.WriteString(fmt.Sprintf("\t%s = 0x8000 + iota - _NUM_NONTERMINALS\n", token))
	}
	file.WriteString(")\n\n")

	file.WriteString("func tokenValue(token uint16) uint16 {\n")
	file.WriteString("\treturn 0x7FFF & token\n")
	file.WriteString("}\n\n")

	file.WriteString("func isTerminal(token uint16) bool {\n")
	file.WriteString("\treturn token >= 0x800\n")
	file.WriteString("}\n\n")

	file.WriteString("func tokenToString(token uint16) string {\n")
	file.WriteString("\tswitch token {\n")
	for _, token := range nonterminals {
		file.WriteString(fmt.Sprintf("\tcase %s:\n", token))
		file.WriteString(fmt.Sprintf("\t\treturn \"%s\"\n", token))
	}
	for _, token := range terminals {
		file.WriteString(fmt.Sprintf("\tcase %s:\n", token))
		file.WriteString(fmt.Sprintf("\t\treturn \"%s\"\n", token))
	}
	file.WriteString("\t}\n")
	file.WriteString("\treturn \"UNKNOWN_TOKEN\"\n")
	file.WriteString("}")

	return nil
}

func emitRules(outdir string, rules []rule) error {
	outPath := outdir + "/" + "rules.go"
	file, err := createFile(outPath)

	if err != nil {
		return err
	}

	defer file.Close()

	maxRHSLen := 0

	for _, rule := range rules {
		if len(rule.RHS) > maxRHSLen {
			maxRHSLen = len(rule.RHS)
		}
	}

	packageName := path.Base(outdir)

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	file.WriteString("import (\n")
	file.WriteString("\t\"errors\"\n")
	file.WriteString("\t\"fmt\"\n")
	file.WriteString(")\n\n")

	file.WriteString("/*\n")
	file.WriteString("rule represents a grammar rule of the language. Its lhs is a single token while its rhs is a slice of tokens.\n")
	file.WriteString("*/\n")
	file.WriteString("type rule struct {\n")
	file.WriteString("\tlhs uint16\n")
	file.WriteString("\trhs []uint16\n")
	file.WriteString("}\n\n")

	file.WriteString("/*\n")
	file.WriteString("The maximum length of the rhs of a rule of the language\n")
	file.WriteString("*/\n")
	file.WriteString(fmt.Sprintf("const _MAX_RHS_LEN = %d\n\n", maxRHSLen))

	file.WriteString("/*\n")
	file.WriteString("The rules of the language. They are sorted by their rhs\n")
	file.WriteString("*/\n")
	file.WriteString("var _RULES = []rule{\n")
	for _, rule := range rules {
		file.WriteString(fmt.Sprintf("\trule{%s, []uint16{%s}},\n", rule.LHS, strings.Join(rule.RHS, ", ")))
	}
	file.WriteString("}\n\n")

	file.WriteString(
		`/*
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
}`)

	return nil
}

func emitFunction(outdir string, preamble string, rules []rule) error {
	outPath := outdir + "/" + "function.go"
	file, err := createFile(outPath)

	if err != nil {
		return err
	}

	defer file.Close()

	packageName := path.Base(outdir)

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	file.WriteString(preamble)
	file.WriteString("\n\n")

	file.WriteString("/*\n")
	file.WriteString("function is the semantic function of the parser.\n")
	file.WriteString("*/\n")
	file.WriteString("func function(thread int, ruleNum int, lhs *symbol, rhs []*symbol) {\n")
	file.WriteString("\tswitch ruleNum {\n")
	for i, rule := range rules {
		file.WriteString(fmt.Sprintf("\tcase %d:\n", i))
		file.WriteString(fmt.Sprintf("\t\t%s0 := lhs\n", rule.LHS))
		for j, _ := range rule.RHS {
			file.WriteString(fmt.Sprintf("\t\t%s%d := rhs[%d]\n", rule.RHS[j], j+1, j))
		}
		file.WriteString("\n")
		if len(rule.RHS) > 0 {
			file.WriteString(fmt.Sprintf("\t\t%s0.Child = %s1\n", rule.LHS, rule.RHS[0]))
			for j := 0; j < len(rule.RHS)-1; j++ {
				file.WriteString(fmt.Sprintf("\t\t%s%d.Next = %s%d\n", rule.RHS[j], j+1, rule.RHS[j+1], j+2))
			}
		}
		file.WriteString("\n")
		action := rule.Action
		action = strings.Replace(action, "$$", rule.LHS+"0", -1)
		for j, _ := range rule.RHS {
			action = strings.Replace(action, fmt.Sprintf("$%d", j+1), fmt.Sprintf("%s%d", rule.RHS[j], j+1), -1)
		}
		lines := strings.Split(action, "\n")
		for _, line := range lines {
			file.WriteString("\t\t")
			file.WriteString(line)
			file.WriteString("\n")
		}
	}
	file.WriteString("\t}\n")
	file.WriteString("}\n")

	return nil
}

func emitPrecMatrix(outdir string, terminals stringSet, matrix precMatrix) error {
	outPath := outdir + "/" + "matrix.go"
	file, err := createFile(outPath)

	if err != nil {
		return err
	}

	defer file.Close()

	packageName := path.Base(outdir)

	file.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	file.WriteString("const _YI = _YIELDS_PREC\n")
	file.WriteString("const _EQ = _EQ_PREC\n")
	file.WriteString("const _TA = _TAKES_PREC\n")
	file.WriteString("const _NO = _NO_PREC\n\n")

	pureMatrix := make([][]uint16, len(terminals))

	for i, terminal1 := range terminals {
		pureMatrix[i] = make([]uint16, len(terminals))
		for j, terminal2 := range terminals {
			pureMatrix[i][j] = matrix[terminal1][terminal2]
		}
	}

	file.WriteString("/*\n")
	file.WriteString("The unpacked precedence matrix\n")
	file.WriteString("*/\n")
	file.WriteString("var _PREC_MATRIX [][]uint16 = [][]uint16{\n")
	for i, _ := range terminals {
		file.WriteString("\t[]uint16{")
		for j, _ := range terminals {
			if j < len(terminals)-1 {
				file.WriteString(fmt.Sprintf("%s, ", precToString(pureMatrix[i][j])))
			} else {
				file.WriteString(precToString(pureMatrix[i][j]))
			}
		}
		file.WriteString("},\n")
	}
	file.WriteString("}\n")

	bitPackedMatrix := bitPack(pureMatrix)
	file.WriteString("/*\n")
	file.WriteString("The packed precedence matrix\n")
	file.WriteString("*/\n")
	file.WriteString("var _PREC_MATRIX_BITPACKED []uint64 = []uint64{\n\t")
	for _, v := range bitPackedMatrix {
		file.WriteString(fmt.Sprintf("%d, ", v))
	}
	file.WriteString("\n}\n\n")

	file.WriteString(
		`/*
getPrecedence returns the precedence between two tokens using the bitpacked precedence matrix.
*/
func getPrecedence(token1 uint16, token2 uint16) uint16 {
	tv1 := tokenValue(token1)
	tv2 := tokenValue(token2)

	flatElemPos := tv1*_NUM_TERMINALS + tv2
	elem := _PREC_MATRIX_BITPACKED[flatElemPos/32]
	pos := uint((flatElemPos % 32) * 2)

	return uint16((elem >> pos) & 0x3)
}`)

	return nil
}

/*
bitPack packs the matrix into a slice of uint64 where a precedence value is represented by just 2 bits.
*/
func bitPack(matrix [][]uint16) []uint64 {
	newSize := int(math.Ceil(float64((len(matrix) * len(matrix))) / float64(32)))

	newMatrix := make([]uint64, newSize)

	setPrec := func(elem *uint64, pos uint, prec uint16) {
		bitMask := uint64(0x3 << pos)
		*elem = (*elem & ^bitMask) | (uint64(prec) << pos)
	}

	for i, _ := range matrix {
		for j, prec := range matrix[i] {
			flatElemPos := i*len(matrix) + j
			newElemPtr := &newMatrix[flatElemPos/32]
			newElemPos := uint((flatElemPos % 32) * 2)
			setPrec(newElemPtr, newElemPos, prec)
		}
	}

	return newMatrix
}

func emitCommonFiles(outdir string) error {
	dir := "common"
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return err
	}

	for _, fileInfo := range files {
		if !fileInfo.IsDir() {
			inPath := dir + "/" + fileInfo.Name()
			outPath := outdir + "/" + fileInfo.Name()

			inFileContent, err := ioutil.ReadFile(inPath)

			if err != nil {
				return err
			}

			outFile, err := createFile(outPath)

			if err != nil {
				return err
			}

			packageName := path.Base(outdir)

			outFile.WriteString(fmt.Sprintf("package %s\n\n", packageName))
			outFile.Write(inFileContent)

			outFile.Close()
		}
	}

	return nil
}

func createFile(path string) (*os.File, error) {
	fileExisted := fileExists(path)
	file, err := os.Create(path)

	if err == nil {
		if fileExisted {
			fmt.Printf("Overwritten file %s\n", path)
		} else {
			fmt.Printf("Created file %s\n", path)
		}
	}

	return file, err
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
