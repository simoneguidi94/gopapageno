package generator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseGrammar(filename string) (string, string, []rule) {
	fmt.Println("Specified parser file:", filename)

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	separatorRegex, err := regexp.Compile("^%%\\s*$")
	checkRegexpCompileError(err)
	axiomRegex, err := regexp.Compile("^%axiom\\s*([a-zA-Z][a-zA-Z0-9]*)\\s*$")
	checkRegexpCompileError(err)

	scanner := bufio.NewScanner(file)

	//Scan the preamble
	goPreamble := make([]string, 0)

	for scanner.Scan() {
		curLine := scanner.Text()
		if separatorRegex.MatchString(curLine) {
			break
		}
		goPreamble = append(goPreamble, curLine)
	}

	//Scan the axiom
	axiom := ""
	moreThanOneAxiomWarning := false

	for scanner.Scan() {
		curLine := scanner.Text()
		if separatorRegex.MatchString(curLine) {
			break
		}
		axiomMatch := axiomRegex.FindStringSubmatch(curLine)
		if axiomMatch != nil {
			if axiom != "" && !moreThanOneAxiomWarning {
				fmt.Println("Warning: axiom is defined more than once")
				moreThanOneAxiomWarning = true
			}
			axiom = axiomMatch[1]
		}
	}

	ruleLines := make([]string, 0)
	for scanner.Scan() {
		curLine := scanner.Text()
		ruleLines = append(ruleLines, curLine)
	}

	rules := parseRules(strings.Join(ruleLines, "\n"))

	return strings.Join(goPreamble, "\n"), axiom, rules
}

func parseRules(input string) []rule {
	bytes := []byte(input)

	rules := make([]rule, 0)

	pos := 0

	pos = skipSpaces(bytes, pos)

	for pos < len(bytes) {
		firstRule := rule{}

		var lhs string
		lhs, pos = getIdentifier(bytes, pos)

		if lhs == "" {
			panic("Missing or invalid identifier for lhs")
		}

		firstRule.LHS = lhs

		pos = skipSpaces(bytes, pos)

		if bytes[pos] == ':' {
			pos++
		} else {
			panic("One of the rules is missing a colon between lhs and rhs")
		}

		pos = skipSpaces(bytes, pos)

		firstRule.RHS = make([]string, 0)

		for bytes[pos] != '{' {
			var rhsToken string
			rhsToken, pos = getIdentifier(bytes, pos)
			if rhsToken == "" {
				panic("Invalid identifier for rhs")
			}
			firstRule.RHS = append(firstRule.RHS, rhsToken)
			pos = skipSpaces(bytes, pos)
		}

		var semFun string
		semFun, pos = getSemanticFunction(bytes, pos)

		firstRule.Action = semFun

		rules = append(rules, firstRule)

		for {
			pos = skipSpaces(bytes, pos)
			if bytes[pos] == ';' {
				pos++
				break
			} else if bytes[pos] == '|' {
				pos++

				pos = skipSpaces(bytes, pos)

				nextRule := rule{}
				nextRule.LHS = lhs
				nextRule.RHS = make([]string, 0)

				for bytes[pos] != '{' {
					var rhsToken string
					rhsToken, pos = getIdentifier(bytes, pos)
					if rhsToken == "" {
						panic("Invalid identifier for rhs")
					}
					nextRule.RHS = append(nextRule.RHS, rhsToken)
					pos = skipSpaces(bytes, pos)
				}

				var semFun string
				semFun, pos = getSemanticFunction(bytes, pos)

				nextRule.Action = semFun

				rules = append(rules, nextRule)
			} else {
				panic("Invalid character at the end of a rule")
			}
		}

		pos = skipSpaces(bytes, pos)
	}

	return rules
}
