package generator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseLexer(filename string) ([]lexRule, string) {
	fmt.Println("Specified lexer file:", filename)

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	separatorRegex, err := regexp.Compile("^%%\\s*$")
	checkRegexpCompileError(err)
	definitionRegex, err := regexp.Compile("^([a-zA-Z][a-zA-Z0-9]*)\\s*(.+)$")
	checkRegexpCompileError(err)

	scanner := bufio.NewScanner(file)

	//Scan the definitions section
	definitions := make(map[string]string)

	for scanner.Scan() {
		curLine := scanner.Text()
		if separatorRegex.MatchString(curLine) {
			break
		}
		defMatch := definitionRegex.FindStringSubmatch(curLine)
		if defMatch != nil {
			definitions[defMatch[1]] = strings.TrimSpace(defMatch[2])
		}
	}

	fmt.Println("Definitions:")
	for k, v := range definitions {
		fmt.Printf("%s: %s\n", k, v)
	}

	ruleLines := make([]string, 0)
	for scanner.Scan() {
		curLine := scanner.Text()
		if separatorRegex.MatchString(curLine) {
			break
		}
		ruleLines = append(ruleLines, curLine)
	}

	lexRules := parseLexRules(strings.Join(ruleLines, "\n"), definitions)

	codeLines := make([]string, 0)
	for scanner.Scan() {
		curLine := scanner.Text()
		codeLines = append(codeLines, curLine)
	}

	return lexRules, strings.Join(codeLines, "\n")
}

func parseLexRules(input string, definitions map[string]string) []lexRule {
	bytes := []byte(input)

	lexRules := make([]lexRule, 0)

	pos := 0

	pos = skipSpaces(bytes, pos)

	curRegex := ""

	for pos < len(bytes) {
		startingPos := pos

		//Read anything until a { is reached
		for pos < len(bytes) && bytes[pos] != '{' {
			pos++
		}

		if pos >= len(bytes) {
			break
		}

		//When a { is reached, try to read a definition and then a }
		//If it's not possible then the regex part is over
		curlyLParPos := pos

		pos++

		var identifier string
		identifier, pos = getIdentifier(bytes, pos)
		curlyRParPos := pos
		if bytes[curlyRParPos] == '}' && identifier != "" {
			foundInDefinitions := false
			for key, value := range definitions {
				if key == identifier {
					foundInDefinitions = true
					curRegex += string(bytes[startingPos:curlyLParPos]) + value
					pos++
					break
				}
			}

			if !foundInDefinitions {
				fmt.Println(lexRules)
				panic(fmt.Sprintf("Missing definition \"%s\"", identifier))
			}
		} else {
			pos = curlyLParPos

			var semFun string
			semFun, pos = getSemanticFunction(bytes, pos)

			curRegex += string(bytes[startingPos:curlyLParPos])

			curRegex = strings.Trim(curRegex, " \t\r\n")

			lexRules = append(lexRules, lexRule{curRegex, semFun})

			curRegex = ""

			pos = skipSpaces(bytes, pos)
		}
	}

	return lexRules
}
