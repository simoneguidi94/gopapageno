package generator

import (
	"fmt"
)

type lexRule struct {
	Regex  string
	Action string
}

func (r lexRule) String() string {
	return fmt.Sprintf("%s: %s", r.Regex, r.Action)
}
