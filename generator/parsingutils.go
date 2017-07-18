package generator

import "fmt"

func skipSpaces(bytes []byte, curPos int) int {
	for curPos < len(bytes) && (bytes[curPos] == ' ' || bytes[curPos] == '\t' || bytes[curPos] == '\r' || bytes[curPos] == '\n') {
		curPos++
	}
	return curPos
}

func getIdentifier(bytes []byte, curPos int) (string, int) {
	startingPos := curPos
	if curPos < len(bytes) && ((bytes[curPos] >= 'a' && bytes[curPos] <= 'z') || (bytes[curPos] >= 'A' && bytes[curPos] <= 'Z') || (bytes[curPos] == '_')) {
		curPos++
		for curPos < len(bytes) && ((bytes[curPos] >= 'a' && bytes[curPos] <= 'z') || (bytes[curPos] >= 'A' && bytes[curPos] <= 'Z') || (bytes[curPos] == '_') || (bytes[curPos] >= '9' && bytes[curPos] <= '0')) {
			curPos++
		}
	}
	return string(bytes[startingPos:curPos]), curPos
}

func getSemanticFunction(bytes []byte, curPos int) (string, int) {
	startingPos := curPos
	numBraces := 0
	for curPos < len(bytes) {
		if bytes[curPos] == '\'' {
			curPos++
			escape := false
			for curPos < len(bytes) {
				if bytes[curPos] == '\\' {
					escape = !escape
				} else if bytes[curPos] == '\'' {
					if !escape {
						break
					}
					escape = false
				} else {
					escape = false
				}
				curPos++
			}
		} else if bytes[curPos] == '"' {
			curPos++
			escape := false
			for curPos < len(bytes) {
				if bytes[curPos] == '\\' {
					escape = !escape
				} else if bytes[curPos] == '"' {
					if !escape {
						break
					}
					escape = false
				} else {
					escape = false
				}
				curPos++
			}
		} else if bytes[curPos] == '`' {
			curPos++
			for curPos < len(bytes) {
				if bytes[curPos] == '`' {
					break
				}
				curPos++
			}
		} else if bytes[curPos] == '/' && curPos < len(bytes)-1 {
			curPos++
			if bytes[curPos] == '*' {
				curPos++
				foundStar := false
				for curPos < len(bytes) {
					if bytes[curPos] == '*' {
						foundStar = true
					} else if bytes[curPos] == '/' {
						if foundStar {
							break
						}
						foundStar = false
					} else {
						foundStar = false
					}
					curPos++
				}
			} else if bytes[curPos] == '/' {
				curPos++
				for curPos < len(bytes) {
					if bytes[curPos] == '\n' {
						break
					}
					curPos++
				}
			}
		} else if bytes[curPos] == '{' {
			numBraces++
		} else if bytes[curPos] == '}' {
			numBraces--
			if numBraces == 0 {
				curPos++
				break
			}
		}
		curPos++
	}

	return string(bytes[startingPos:curPos]), curPos
}

func checkRegexpCompileError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
