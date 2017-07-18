package arithmetic

type lexerDfaState struct {
	Transitions     [256]int
	IsFinal         bool
	AssociatedRules []int
}

type lexerDfa []lexerDfaState