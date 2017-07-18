package generator

import (
	"errors"
	"fmt"
)

type precMatrix map[string]map[string]uint16

func newPrecMatrix(terminals stringSet) precMatrix {
	precMatrix := make(map[string]map[string]uint16)

	for _, terminal := range terminals {
		precMatrix[terminal] = make(map[string]uint16)
		for _, terminal2 := range terminals {
			precMatrix[terminal][terminal2] = _NO
		}
	}

	return precMatrix
}

func (matrix precMatrix) String() string {
	s := ""
	for key, row := range matrix {
		s += key + ": [ "
		for key2, prec := range row {
			s += key2 + ":" + precToString(prec) + " "
		}
		s += "]\n"
	}
	bytes := []byte(s)
	return string(bytes[:len(bytes)-1])
}

func createPrecMatrix(rules []rule, nonterminals stringSet, terminals stringSet) (precMatrix, error) {
	precMatrix := newPrecMatrix(terminals)
	lts, rts := getTerminalSets(rules, nonterminals, terminals)

	fmt.Println("lts:", lts)
	fmt.Println("rts:", rts)

	for _, rule := range rules {
		rhs := rule.RHS
		//Check digrams
		for i := 0; i < len(rhs)-1; i++ {
			token1 := rhs[i]
			token2 := rhs[i+1]

			if terminals.Contains(token1) && terminals.Contains(token2) {
				//Check if the matrix already contains an entry for this couple
				if precMatrix[token1][token2] != _NO && precMatrix[token1][token2] != _EQ {
					return precMatrix, errors.New(fmt.Sprintf("Error: the precedence relation is not unique between %s and %s", token1, token2))
				}
				precMatrix[token1][token2] = _EQ
			} else if nonterminals.Contains(token1) && terminals.Contains(token2) {
				for _, token := range *rts[token1] {
					//Check if the matrix already contains an entry for this couple
					if precMatrix[token][token2] != _NO && precMatrix[token][token2] != _TA {
						return precMatrix, errors.New(fmt.Sprintf("Error: the precedence relation is not unique between %s and %s", token, token2))
					}
					precMatrix[token][token2] = _TA
				}
			} else if terminals.Contains(token1) && nonterminals.Contains(token2) {
				for _, token := range *lts[token2] {
					//Check if the matrix already contains an entry for this couple
					if precMatrix[token1][token] != _NO && precMatrix[token1][token] != _YI {
						return precMatrix, errors.New(fmt.Sprintf("Error: the precedence relation is not unique between %s and %s", token1, token))
					}
					precMatrix[token1][token] = _YI
				}
			} else {
				return precMatrix, errors.New(fmt.Sprintf("Error: the rule %s is not in operator precedence form", rule.String()))
			}
		}

		//Check trigrams
		for i := 0; i < len(rhs)-2; i++ {
			token1 := rhs[i]
			token2 := rhs[i+1]
			token3 := rhs[i+2]

			if terminals.Contains(token1) && nonterminals.Contains(token2) && terminals.Contains(token3) {
				//Check if the matrix already contains an entry for this couple
				if precMatrix[token1][token3] != _NO && precMatrix[token1][token3] != _EQ {
					return precMatrix, errors.New(fmt.Sprintf("Error: the precedence relation is not unique between %s and %s", token1, token3))
				}
				precMatrix[token1][token3] = _EQ
			}
		}
	}
	//Set precedence for #
	for _, terminal := range terminals {
		if terminal != "_TERM" {
			precMatrix["_TERM"][terminal] = _YI
			precMatrix[terminal]["_TERM"] = _TA
		}
	}
	precMatrix["_TERM"]["_TERM"] = _EQ

	return precMatrix, nil
}

func getTerminalSets(rules []rule, nonterminals stringSet, terminals stringSet) (map[string]*stringSet, map[string]*stringSet) {
	lts := make(map[string]*stringSet)
	rts := make(map[string]*stringSet)

	for _, nonterminal := range nonterminals {
		ltsSet := newStringSet()
		lts[nonterminal] = &ltsSet
		rtsSet := newStringSet()
		rts[nonterminal] = &rtsSet
	}

	//Add direct terminals
	for _, rule := range rules {
		for i := 0; i < len(rule.RHS); i++ {
			token := rule.RHS[i]
			if terminals.Contains(token) {
				lts[rule.LHS].Add(token)
				break
			}
		}

		for i := len(rule.RHS) - 1; i >= 0; i-- {
			token := rule.RHS[i]
			if terminals.Contains(token) {
				rts[rule.LHS].Add(token)
				break
			}
		}
	}

	//Add indirect terminals
	modified := true
	for modified {
		modified = false
		for _, rule := range rules {
			lhs := rule.LHS
			rhs := rule.RHS

			firstToken := rhs[0]
			if nonterminals.Contains(firstToken) {
				for _, token := range *lts[firstToken] {
					if !lts[lhs].Contains(token) {
						lts[lhs].Add(token)
						modified = true
					}
				}
			}

			lastToken := rhs[len(rhs)-1]
			if nonterminals.Contains(lastToken) {
				for _, token := range *rts[lastToken] {
					if !rts[lhs].Contains(token) {
						rts[lhs].Add(token)
						modified = true
					}
				}
			}
		}
	}

	return lts, rts
}
