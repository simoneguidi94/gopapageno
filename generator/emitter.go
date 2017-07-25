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

func emitRules(outdir string, rules []rule, nonterminals stringSet, terminals stringSet) error {
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

	trie := createTrie(rules, nonterminals, terminals)
	compressedTrie := trie.Compress(nonterminals, terminals)

	file.WriteString("var compressedTrie = []uint16{")
	if len(compressedTrie) > 0 {
		file.WriteString(fmt.Sprintf("%d", compressedTrie[0]))
		for i := 1; i < len(compressedTrie); i++ {
			file.WriteString(fmt.Sprintf(", %d", compressedTrie[i]))
		}
	}
	file.WriteString("}\n\n")

	file.WriteString(
		`/*
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
	file.WriteString("func function(thread int, ruleNum uint16, lhs *symbol, rhs []*symbol) {\n")
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
