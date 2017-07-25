package main

import (
	"github.com/simoneguidi94/gopapageno/generator"
)

func main() {
	//generator.Generate("languages/arithmetic/lexer/arith.l", "languages/arithmetic/parser/arith.g", "languages/arithmetic")
	generator.Generate("languages/xml/lexer/xml.l", "languages/xml/parser/xml.g", "languages/xml")
}
