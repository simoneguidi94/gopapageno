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
