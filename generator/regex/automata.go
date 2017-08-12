package regex

const _EPSILON = 0

type NfaState struct {
	Transitions     [256][]*NfaState
	AssociatedRules []int
}

func (state *NfaState) EpsilonClosure() []*NfaState {
	closure := make([]*NfaState, 0)

	closure = append(closure, state)

	nextStateToCheckPos := 0

	for nextStateToCheckPos < len(closure) {
		curState := closure[nextStateToCheckPos]
		closure = append(closure, curState.Transitions[_EPSILON]...)
		nextStateToCheckPos++
	}

	return closure
}

func stateSetContains(stateSet []*NfaState, state *NfaState) bool {
	for _, curState := range stateSet {
		if curState == state {
			return true
		}
	}
	return false
}

func stateSetEpsilonClosure(stateSet []*NfaState) []*NfaState {
	closure := make([]*NfaState, 0, len(stateSet))
	for _, curState := range stateSet {
		curEpsilonClosure := curState.EpsilonClosure()

		for _, curClosureState := range curEpsilonClosure {
			if !stateSetContains(closure, curClosureState) {
				closure = append(closure, curClosureState)
			}
		}
	}

	return closure
}

func stateSetMove(stateSet []*NfaState, char int) []*NfaState {
	states := make([]*NfaState, 0)

	for _, curState := range stateSet {
		reachedStates := curState.Transitions[char]

		for _, curReachedState := range reachedStates {
			if !stateSetContains(states, curReachedState) {
				states = append(states, curReachedState)
			}
		}
	}

	return states
}

func (state *NfaState) AddTransition(char byte, ptr *NfaState) {
	if state.Transitions[char] == nil {
		state.Transitions[char] = []*NfaState{ptr}
	} else {
		state.Transitions[char] = append(state.Transitions[char], ptr)
	}
}

type Nfa struct {
	Initial   *NfaState
	Final     *NfaState
	NumStates int
}

func NewEmptyStringNfa() Nfa {
	nfa := Nfa{}
	nfaInitialFinal := NfaState{}

	nfa.Initial = &nfaInitialFinal
	nfa.Final = &nfaInitialFinal

	nfa.NumStates = 1

	return nfa
}

func newNfaFromChar(char byte) Nfa {
	nfa := Nfa{}
	nfaInitial := NfaState{}
	nfaFinal := NfaState{}

	nfaInitial.AddTransition(char, &nfaFinal)
	nfa.Initial = &nfaInitial
	nfa.Final = &nfaFinal

	nfa.NumStates = 2

	return nfa
}

func newNfaFromCharClass(chars [256]bool) Nfa {
	nfa := Nfa{}
	nfaInitial := NfaState{}
	nfaFinal := NfaState{}

	for i, thereIs := range chars {
		if thereIs {
			nfaInitial.AddTransition(byte(i), &nfaFinal)
		}
	}

	nfa.Initial = &nfaInitial
	nfa.Final = &nfaFinal

	nfa.NumStates = 2

	return nfa
}

func newNfaFromString(str []byte) Nfa {
	nfa := Nfa{}
	firstState := NfaState{}

	curState := &firstState

	nfa.Initial = curState

	for _, curChar := range str {
		newState := NfaState{}
		curState.AddTransition(curChar, &newState)
		curState = &newState
	}

	nfa.Final = curState

	nfa.NumStates = len(str) + 1

	return nfa
}

func (nfa1 *Nfa) Concatenate(nfa2 Nfa) {
	*nfa1.Final = *nfa2.Initial
	nfa1.Final = nfa2.Final

	nfa1.NumStates = nfa1.NumStates + nfa2.NumStates - 1
}

//Operator |
func (nfa1 *Nfa) Unite(nfa2 Nfa) {
	newInitial := NfaState{}
	newFinal := NfaState{}

	newInitial.AddTransition(_EPSILON, nfa1.Initial)
	newInitial.AddTransition(_EPSILON, nfa2.Initial)

	nfa1.Final.AddTransition(_EPSILON, &newFinal)
	nfa2.Final.AddTransition(_EPSILON, &newFinal)

	nfa1.Initial = &newInitial
	nfa1.Final = &newFinal

	nfa1.NumStates += nfa2.NumStates + 2
}

//Operator *
func (nfa *Nfa) KleeneStar() {
	newInitial := NfaState{}
	newFinal := NfaState{}

	newInitial.AddTransition(_EPSILON, nfa.Initial)
	newInitial.AddTransition(_EPSILON, &newFinal)

	nfa.Final.AddTransition(_EPSILON, nfa.Initial)
	nfa.Final.AddTransition(_EPSILON, &newFinal)

	nfa.Initial = &newInitial
	nfa.Final = &newFinal

	nfa.NumStates += 2
}

//Operator +
func (nfa *Nfa) KleenePlus() {
	newInitial := NfaState{}
	newFinal := NfaState{}

	newInitial.AddTransition(_EPSILON, nfa.Initial)

	nfa.Final.AddTransition(_EPSILON, nfa.Initial)
	nfa.Final.AddTransition(_EPSILON, &newFinal)

	nfa.Initial = &newInitial
	nfa.Final = &newFinal

	nfa.NumStates += 2
}

//Operator ?
func (nfa *Nfa) ZeroOrOne() {
	newInitial := NfaState{}
	newFinal := NfaState{}

	newInitial.AddTransition(_EPSILON, nfa.Initial)
	newInitial.AddTransition(_EPSILON, &newFinal)

	nfa.Final.AddTransition(_EPSILON, &newFinal)

	nfa.Initial = &newInitial
	nfa.Final = &newFinal

	nfa.NumStates += 2
}

func (nfa *Nfa) AddAssociatedRule(ruleNum int) {
	finalState := nfa.Final

	if finalState.AssociatedRules == nil {
		finalState.AssociatedRules = []int{ruleNum}
	} else {
		finalState.AssociatedRules = append(finalState.AssociatedRules, ruleNum)
	}
}

func (nfa *Nfa) ToDfa() Dfa {
	genStates := make([]nfaStateSetPtr, 0)

	curDfaStateNum := 0

	initialDfaState := DfaState{}
	initialDfaState.Num = curDfaStateNum

	dfa := Dfa{&initialDfaState, make([]*DfaState, 0), 1}

	genStates = append(genStates, nfaStateSetPtr{nfa.Initial.EpsilonClosure(), &initialDfaState})

	search := func(gStates []nfaStateSetPtr, stateSet []*NfaState) *nfaStateSetPtr {
		for _, curGState := range gStates {
			if len(curGState.StateSet) != len(stateSet) {
				continue
			}
			equal := true
			for i, _ := range curGState.StateSet {
				if curGState.StateSet[i] != stateSet[i] {
					equal = false
					break
				}
			}
			if equal {
				return &curGState
			}
		}
		return nil
	}

	nextStateToCheckPos := 0

	for nextStateToCheckPos < len(genStates) {
		curStateSet := genStates[nextStateToCheckPos].StateSet
		curDfaState := genStates[nextStateToCheckPos].Ptr

		//For each character
		for i := 1; i < 256; i++ {
			charStateSet := stateSetMove(curStateSet, i)
			epsilonClosure := stateSetEpsilonClosure(charStateSet)

			if len(epsilonClosure) == 0 {
				continue
			}

			foundStateSetPtr := search(genStates, epsilonClosure)

			if foundStateSetPtr != nil {
				curDfaState.Transitions[i] = foundStateSetPtr.Ptr
			} else {
				curDfaStateNum++
				newDfaState := DfaState{}
				newDfaState.Num = curDfaStateNum
				newDfaState.AssociatedRules = make([]int, 0)
				for _, curNfaState := range epsilonClosure {
					newDfaState.AssociatedRules = append(newDfaState.AssociatedRules, curNfaState.AssociatedRules...)
				}
				curDfaState.Transitions[i] = &newDfaState
				newStateSetPtr := nfaStateSetPtr{epsilonClosure, &newDfaState}

				genStates = append(genStates, newStateSetPtr)
			}
		}
		nextStateToCheckPos++
	}

	dfa.NumStates = len(genStates)

	for _, genState := range genStates {
		if stateSetContains(genState.StateSet, nfa.Final) {
			finalState := genState.Ptr
			finalState.IsFinal = true
			dfa.Final = append(dfa.Final, finalState)
		}
	}

	return dfa
}

type DfaState struct {
	Num             int
	Transitions     [256]*DfaState
	IsFinal         bool
	AssociatedRules []int
}

type Dfa struct {
	Initial   *DfaState
	Final     []*DfaState
	NumStates int
}

func (dfaState *DfaState) getStatesR(addedStates *[]*DfaState) {
	//The state was already added, return
	if (*addedStates)[dfaState.Num] != nil {
		return
	}
	(*addedStates)[dfaState.Num] = dfaState

	for _, nextState := range dfaState.Transitions {
		if nextState != nil {
			nextState.getStatesR(addedStates)
		}
	}
}

/*
GetState returns a slice containing all the states of the dfa.
The states are sorted by their state number.
*/
func (dfa *Dfa) GetStates() []*DfaState {
	states := make([]*DfaState, dfa.NumStates)

	dfa.Initial.getStatesR(&states)

	return states
}

/*func (dfa *Dfa) Check(str []byte) (bool, bool, uint16) {
	curState := dfa.Initial

	//fmt.Println(curState)

	for _, curChar := range str {
		curState = curState.Transitions[curChar]

		//fmt.Println(curState)

		if curState == nil {
			return false, false, 0
		}
	}

	if len(curState.AssociatedTokens) == 0 {
		return curState.IsFinal, false, 0
	}

	index := 0
	minRule := curState.AssociatedTokens[0].RuleNum

	for i := 1; i < len(curState.AssociatedTokens); i++ {
		if curState.AssociatedTokens[i].RuleNum < minRule {
			minRule = curState.AssociatedTokens[i].RuleNum
			index = i
		}
	}

	return curState.IsFinal, true, curState.AssociatedTokens[index].Token
}*/

type nfaStateSetPtr struct {
	StateSet []*NfaState
	Ptr      *DfaState
}

func (stateSet1 *nfaStateSetPtr) Equals(stateSet2 *nfaStateSetPtr) bool {
	if len(stateSet1.StateSet) != len(stateSet2.StateSet) {
		return false
	}
	for i, _ := range stateSet1.StateSet {
		if stateSet1.StateSet[i] != stateSet2.StateSet[i] {
			return false
		}
	}
	return true
}
