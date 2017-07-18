package generator

import (
	"fmt"
	"strings"
)

func deleteRepeatedRHS(nonterminals stringSet, terminals stringSet, axiom string, rules []rule) ([]rule, stringSet) {
	newRules := make([]rule, 0)

	dictRules := createNewDictRules()

	for i, rule := range rules {
		dictRules.Add(rule.RHS, rule.LHS, &rules[i].Action)
	}

	fmt.Println("Old dictRules:")
	fmt.Println(dictRules)

	updatedDictRules := deleteCopyRules(nonterminals, rules, dictRules)

	fmt.Println("Updated dictRules:")
	fmt.Println(updatedDictRules)

	V := updatedDictRules.ToValueLHSSets()
	//fmt.Println("ValueLHSSets:")
	//fmt.Println(V)

	newDictRules := createNewDictRules()
	copyDictRules := updatedDictRules.Copy()

	for i, _ := range copyDictRules.KeysRHS {
		keyRHS := copyDictRules.KeysRHS[i]
		valueLHS := copyDictRules.ValuesLHS[i]
		semAction := copyDictRules.SemActions[i]

		isTerminalRule := true

		for _, token := range keyRHS {
			if nonterminals.Contains(token) {
				isTerminalRule = false
				break
			}
		}

		if isTerminalRule {
			for _, curLHS := range valueLHS {
				newDictRules.Add(keyRHS, curLHS, semAction)
			}
			updatedDictRules.Remove(keyRHS)
		}
	}

	//fmt.Println("New dictRules:")
	//fmt.Println(newDictRules)

	dictRulesForIteration := createNewDictRules()
	loop := true

	for loop {
		for i, _ := range updatedDictRules.KeysRHS {
			keyRHS := updatedDictRules.KeysRHS[i]
			valueLHS := updatedDictRules.ValuesLHS[i]
			semAction := updatedDictRules.SemActions[i]

			newRuleRHS := make([]string, 0)
			addNewRules(&dictRulesForIteration, keyRHS, valueLHS, semAction, nonterminals, V, newRuleRHS)
		}
		valueLHSSets := dictRulesForIteration.ToValueLHSSets()
		addedNonterminals := make([]stringSet, 0)

		for _, curNontermSet := range valueLHSSets {
			contained := false
			for _, otherNonTermSet := range V {
				if curNontermSet.Equals(otherNonTermSet) {
					contained = true
					break
				}
			}
			if !contained {
				addedNonterminals = append(addedNonterminals, curNontermSet)
				V = append(V, curNontermSet)
			}
		}

		for i, _ := range dictRulesForIteration.KeysRHS {
			keyRHS := dictRulesForIteration.KeysRHS[i]
			valueLHS := dictRulesForIteration.ValuesLHS[i]
			semAction := dictRulesForIteration.SemActions[i]

			for _, curLHS := range valueLHS {
				newDictRules.Add(keyRHS, curLHS, semAction)
			}
		}

		if len(addedNonterminals) == 0 {
			loop = false
		}
	}

	//TODO remove unused nonterminals (see cpapageno)

	newAxiom := "NEW_AXIOM"
	newAxiomSet := newStringSet()
	newAxiomSet.Add(newAxiom)

	newAxiomSemAction := "{\n\t$$.Value = $1.Value\n}"

	V = append(V, newAxiomSet)

	for _, nontermSet := range V {
		if nontermSet.Contains(axiom) {
			newDictRules.Add([]string{strings.Join(nontermSet, "_")}, newAxiom, &newAxiomSemAction)
		}
	}

	fmt.Println("New dictRules:")
	fmt.Println(newDictRules)

	//Create the rules from dictRules
	for i, _ := range newDictRules.KeysRHS {
		keyRHS := newDictRules.KeysRHS[i]
		valueLHS := newDictRules.ValuesLHS[i]
		semAction := newDictRules.SemActions[i]

		newRules = append(newRules, rule{strings.Join(valueLHS, "_"), keyRHS, *semAction})
	}

	newNonterminalSet, _ := inferTokens(newRules)

	return newRules, newNonterminalSet
}

type dictRules struct {
	KeysRHS    [][]string
	ValuesLHS  []stringSet
	SemActions []*string
}

func createNewDictRules() dictRules {
	return dictRules{make([][]string, 0), make([]stringSet, 0), make([]*string, 0)}
}

func (dict *dictRules) Add(keyRHS []string, valueLHS string, semAction *string) {
	found := false
	for i, curKeyRHS := range dict.KeysRHS {
		if rhsEquals(curKeyRHS, keyRHS) {
			dict.ValuesLHS[i].Add(valueLHS)
			found = true
		}
	}

	if !found {
		dict.KeysRHS = append(dict.KeysRHS, keyRHS)
		dict.ValuesLHS = append(dict.ValuesLHS, newStringSet())
		dict.ValuesLHS[len(dict.ValuesLHS)-1].Add(valueLHS)
		dict.SemActions = append(dict.SemActions, semAction)
	}
}

func (dict *dictRules) Remove(keyRHS []string) {
	for i, curKeyRHS := range dict.KeysRHS {
		if rhsEquals(curKeyRHS, keyRHS) {
			dict.KeysRHS = append(dict.KeysRHS[:i], dict.KeysRHS[i+1:]...)
			dict.ValuesLHS = append(dict.ValuesLHS[:i], dict.ValuesLHS[i+1:]...)
			dict.SemActions = append(dict.SemActions[:i], dict.SemActions[i+1:]...)
		}
	}
}

func (dict *dictRules) Copy() dictRules {
	newDict := createNewDictRules()

	for i, _ := range dict.KeysRHS {
		newDict.KeysRHS = append(newDict.KeysRHS, make([]string, len(dict.KeysRHS[i])))
		copy(newDict.KeysRHS[i], dict.KeysRHS[i])
		newDict.ValuesLHS = append(newDict.ValuesLHS, newStringSet())
		for _, curLHS := range dict.ValuesLHS[i] {
			newDict.ValuesLHS[i].Add(curLHS)
		}
		newDict.SemActions = append(newDict.SemActions, dict.SemActions[i])
	}

	return newDict
}

func (dict *dictRules) ToValueLHSSets() []stringSet {
	valueLHSSets := make([]stringSet, 0)

	for _, curLHSSet := range dict.ValuesLHS {
		alreadyContained := false
		for _, curSet := range valueLHSSets {
			if curSet.Equals(curLHSSet) {
				alreadyContained = true
				break
			}
		}
		if !alreadyContained {
			valueLHSSets = append(valueLHSSets, curLHSSet)
		}
	}

	return valueLHSSets
}

func (dict dictRules) String() string {
	s := ""
	for i, _ := range dict.KeysRHS {
		s += fmt.Sprintln(dict.KeysRHS[i], "->", strings.Join(dict.ValuesLHS[i], ", "))
		if dict.SemActions[i] != nil {
			s += fmt.Sprintln("Semantic action:", *dict.SemActions[i])
		} else {
			s += fmt.Sprintln("No semantic action")
		}
	}

	bytes := []byte(s)
	return string(bytes[:len(bytes)-1])
}

func deleteCopyRules(nonterminals stringSet, rules []rule, dictRules dictRules) dictRules {
	copySets := make(map[string]*stringSet)
	rhsDict := make(map[string][][]string)

	copySetsStorage := make([]stringSet, len(nonterminals))

	for i, nonterminal := range nonterminals {
		copySetsStorage[i] = newStringSet()
		copySets[nonterminal] = &copySetsStorage[i]
	}

	for _, rule := range rules {
		if len(rule.RHS) == 1 && nonterminals.Contains(rule.RHS[0]) {
			copySets[rule.LHS].Add(rule.RHS[0])
			dictRules.Remove(rule.RHS)
		} else {
			if _, hasKey := rhsDict[rule.LHS]; hasKey {
				rhsDict[rule.LHS] = append(rhsDict[rule.LHS], rule.RHS)
			} else {
				rhsDict[rule.LHS] = [][]string{rule.RHS}
			}
		}
	}

	changedCopySets := true
	for changedCopySets {
		changedCopySets = false

		for _, nonterminal := range nonterminals {
			lenCopySet := len(*copySets[nonterminal])

			iterCopy := copySets[nonterminal].Copy()

			for _, curCopyRHS := range iterCopy {
				for _, curNonterminal := range *copySets[curCopyRHS] {
					copySets[nonterminal].Add(curNonterminal)
				}
			}

			if lenCopySet < len(*copySets[nonterminal]) {
				changedCopySets = true
			}
		}
	}

	for _, nonterminal := range nonterminals {
		for _, curCopyRHS := range *copySets[nonterminal] {
			rhsDictCopyRHSs := rhsDict[curCopyRHS]
			for _, rhs := range rhsDictCopyRHSs {
				//There's no need to specify semantic actions because they are already linked to the proper rhs
				dictRules.Add(rhs, nonterminal, nil)
			}
		}
	}

	return dictRules
}

func addNewRules(dictRulesForIteration *dictRules, keyRHS []string, valueLHS stringSet, semAction *string, nonterminals stringSet, newNonterminals []stringSet, newRuleRHS []string) {
	if len(keyRHS) == 0 {
		for _, curLHS := range valueLHS {
			dictRulesForIteration.Add(newRuleRHS, curLHS, semAction)
		}
		return
	}
	token := keyRHS[0]
	if nonterminals.Contains(token) {
		for _, nonTermSuperSet := range newNonterminals {
			if nonTermSuperSet.Contains(token) {
				newRuleRHS = append(newRuleRHS, strings.Join(nonTermSuperSet, "_"))
				addNewRules(dictRulesForIteration, keyRHS[1:], valueLHS, semAction, nonterminals, newNonterminals, newRuleRHS)
				newRuleRHSCopy := make([]string, len(newRuleRHS)-1)
				copy(newRuleRHSCopy, newRuleRHS)
				newRuleRHS = newRuleRHSCopy
			}
		}
	} else {
		newRuleRHS = append(newRuleRHS, token)
		addNewRules(dictRulesForIteration, keyRHS[1:], valueLHS, semAction, nonterminals, newNonterminals, newRuleRHS)
		newRuleRHSCopy := make([]string, len(newRuleRHS)-1)
		copy(newRuleRHSCopy, newRuleRHS)
		newRuleRHS = newRuleRHSCopy
	}
}

func rhsEquals(rhs1 []string, rhs2 []string) bool {
	if len(rhs1) != len(rhs2) {
		return false
	}

	for i, _ := range rhs1 {
		if rhs1[i] != rhs2[i] {
			return false
		}
	}

	return true
}

func sortRulesByRHS(rules []rule, nonterminals stringSet, terminals stringSet) []rule {
	sortedRules := make([]rule, 0, len(rules))

	for _, curRule := range rules {
		insertPosition := -1
		for i, curSortedRule := range sortedRules {
			if rhsLessThan(curRule.RHS, curSortedRule.RHS, nonterminals, terminals) {
				insertPosition = i
				break
			}
		}
		if insertPosition == -1 {
			sortedRules = append(sortedRules, curRule)
		} else {
			sortedRules = append(sortedRules, rule{})
			copy(sortedRules[insertPosition+1:], sortedRules[insertPosition:])
			sortedRules[insertPosition] = curRule
		}
	}

	return sortedRules
}

func rhsLessThan(rhs1 []string, rhs2 []string, nonterminals stringSet, terminals stringSet) bool {
	minLen := len(rhs1)
	if len(rhs2) < minLen {
		minLen = len(rhs2)
	}

	for i := 0; i < minLen; i++ {
		//If the first is in nonterminals and the second is in terminals,
		//the first token is certainly less than the second
		if nonterminals.Contains(rhs1[i]) && terminals.Contains(rhs2[i]) {
			return true
		}
		//If the first is in terminals and the second is in nonterminals,
		//the first token is certainly greater than the second
		if terminals.Contains(rhs1[i]) && nonterminals.Contains(rhs2[i]) {
			return false
		}

		if rhs1[i] < rhs2[i] {
			return true
		}
		if rhs1[i] > rhs2[i] {
			return false
		}
	}

	if len(rhs1) < len(rhs2) {
		return true
	}
	return false
}
