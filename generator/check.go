package generator

func inferTokens(rules []rule) (stringSet, stringSet) {
	nonterminals := newStringSet()
	tokens := newStringSet()

	for _, rule := range rules {
		if !nonterminals.Contains(rule.LHS) {
			nonterminals.Add(rule.LHS)
		}
		for _, token := range rule.RHS {
			if !tokens.Contains(token) {
				tokens.Add(token)
			}
		}
	}

	terminals := newStringSet()

	for _, token := range tokens {
		if !nonterminals.Contains(token) {
			terminals.Add(token)
		}
	}

	terminals.Add("_TERM")

	return nonterminals, terminals
}

func checkAxiomUsage(rules []rule, axiom string) bool {
	for _, rule := range rules {
		if rule.LHS == axiom {
			return true
		}
	}
	return false
}
