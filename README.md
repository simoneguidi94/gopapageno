Go PAPAGENO
========

Go PAPAGENO (PArallel PArser GENeratOr) is a parallel parser generator based on Floyd's Operator Precedence Grammars.

It generates parallel Go parsers starting from a lexer and a grammar specification.

The generated parsers are self-contained and can be used without further effort.

This work is based on [Papageno](https://github.com/PAPAGENO-devels/papageno), a C parallel parser generator.

**Note:** at the moment the whole generative part of the project is missing. It is just a parser of arithmetic expressions.

### Installation
```
go get github.com/simoneguidi94/gopapageno
```

### Example

```go
package main

import (
    "fmt"

    "github.com/simoneguidi94/gopapageno/languages/arithmetic"
)

func main() {
    success, result := arithmetic.Parse("expression.txt", 2)
    
    if success {
        fmt.Printf("Result: %d\n", result)
    } else {
        fmt.Printf("Parse failed!")
    }
}
```

### Authors and Contributors

 * Simone Guidi <simone.guidi@mail.polimi.it>