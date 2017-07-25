package regex

/*
parserPreallocMem initializes all the memory pools required by the semantic function.
*/
func parserPreallocMem(inputSize int, numThreads int) {
}

/*
function is the semantic function of the parser.
*/
func function(thread int, ruleNum uint16, lhs *symbol, rhs []*symbol) {
	switch ruleNum {
	case 0:
		NEW_AXIOM0 := lhs
		CONCATENATION_RE_SIMPLE_RE1 := rhs[0]

		NEW_AXIOM0.Child = CONCATENATION_RE_SIMPLE_RE1

		{
			NEW_AXIOM0.Value = CONCATENATION_RE_SIMPLE_RE1.Value
		}
	case 1:
		RE_UNION0 := lhs
		CONCATENATION_RE_SIMPLE_RE1 := rhs[0]
		pipe2 := rhs[1]
		CONCATENATION_RE_SIMPLE_RE3 := rhs[2]

		RE_UNION0.Child = CONCATENATION_RE_SIMPLE_RE1
		CONCATENATION_RE_SIMPLE_RE1.Next = pipe2
		pipe2.Next = CONCATENATION_RE_SIMPLE_RE3

		{
			leftNfa := CONCATENATION_RE_SIMPLE_RE1.Value.(*Nfa)
			rightNfa := CONCATENATION_RE_SIMPLE_RE3.Value.(*Nfa)
			
			leftNfa.Unite(*rightNfa)
			
			RE_UNION0.Value = leftNfa
		}
	case 2:
		RE_UNION0 := lhs
		CONCATENATION_RE_SIMPLE_RE1 := rhs[0]
		pipe2 := rhs[1]
		RE_SIMPLE_RE3 := rhs[2]

		RE_UNION0.Child = CONCATENATION_RE_SIMPLE_RE1
		CONCATENATION_RE_SIMPLE_RE1.Next = pipe2
		pipe2.Next = RE_SIMPLE_RE3

		{
			leftNfa := CONCATENATION_RE_SIMPLE_RE1.Value.(*Nfa)
			rightNfa := RE_SIMPLE_RE3.Value.(*Nfa)
			
			leftNfa.Unite(*rightNfa)
			
			RE_UNION0.Value = leftNfa
		}
	case 3:
		NEW_AXIOM0 := lhs
		RE_SIMPLE_RE1 := rhs[0]

		NEW_AXIOM0.Child = RE_SIMPLE_RE1

		{
			NEW_AXIOM0.Value = RE_SIMPLE_RE1.Value
		}
	case 4:
		RE_UNION0 := lhs
		RE_SIMPLE_RE1 := rhs[0]
		pipe2 := rhs[1]
		CONCATENATION_RE_SIMPLE_RE3 := rhs[2]

		RE_UNION0.Child = RE_SIMPLE_RE1
		RE_SIMPLE_RE1.Next = pipe2
		pipe2.Next = CONCATENATION_RE_SIMPLE_RE3

		{
			leftNfa := RE_SIMPLE_RE1.Value.(*Nfa)
			rightNfa := CONCATENATION_RE_SIMPLE_RE3.Value.(*Nfa)
			
			leftNfa.Unite(*rightNfa)
			
			RE_UNION0.Value = leftNfa
		}
	case 5:
		RE_UNION0 := lhs
		RE_SIMPLE_RE1 := rhs[0]
		pipe2 := rhs[1]
		RE_SIMPLE_RE3 := rhs[2]

		RE_UNION0.Child = RE_SIMPLE_RE1
		RE_SIMPLE_RE1.Next = pipe2
		pipe2.Next = RE_SIMPLE_RE3

		{
			leftNfa := RE_SIMPLE_RE1.Value.(*Nfa)
			rightNfa := RE_SIMPLE_RE3.Value.(*Nfa)
			
			leftNfa.Unite(*rightNfa)
			
			RE_UNION0.Value = leftNfa
		}
	case 6:
		NEW_AXIOM0 := lhs
		RE_UNION1 := rhs[0]

		NEW_AXIOM0.Child = RE_UNION1

		{
			NEW_AXIOM0.Value = RE_UNION1.Value
		}
	case 7:
		RE_UNION0 := lhs
		RE_UNION1 := rhs[0]
		pipe2 := rhs[1]
		CONCATENATION_RE_SIMPLE_RE3 := rhs[2]

		RE_UNION0.Child = RE_UNION1
		RE_UNION1.Next = pipe2
		pipe2.Next = CONCATENATION_RE_SIMPLE_RE3

		{
			leftNfa := RE_UNION1.Value.(*Nfa)
			rightNfa := CONCATENATION_RE_SIMPLE_RE3.Value.(*Nfa)
			
			leftNfa.Unite(*rightNfa)
			
			RE_UNION0.Value = leftNfa
		}
	case 8:
		RE_UNION0 := lhs
		RE_UNION1 := rhs[0]
		pipe2 := rhs[1]
		RE_SIMPLE_RE3 := rhs[2]

		RE_UNION0.Child = RE_UNION1
		RE_UNION1.Next = pipe2
		pipe2.Next = RE_SIMPLE_RE3

		{
			leftNfa := RE_UNION1.Value.(*Nfa)
			rightNfa := RE_SIMPLE_RE3.Value.(*Nfa)
			
			leftNfa.Unite(*rightNfa)
			
			RE_UNION0.Value = leftNfa
		}
	case 9:
		RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]

		RE_SIMPLE_RE0.Child = any1

		{
			RE_SIMPLE_RE0.Value = any1.Value
		}
	case 10:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]

		CONCATENATION_RE_SIMPLE_RE0.Child = any1
		any1.Next = CONCATENATION_RE_SIMPLE_RE2

		{
			leftNfa := any1.Value.(*Nfa)
			rightNfa := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			
			leftNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = leftNfa
		}
	case 11:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]

		CONCATENATION_RE_SIMPLE_RE0.Child = any1
		any1.Next = RE_SIMPLE_RE2

		{
			leftNfa := any1.Value.(*Nfa)
			rightNfa := RE_SIMPLE_RE2.Value.(*Nfa)
			
			leftNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = leftNfa
		}
	case 12:
		RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		plus2 := rhs[1]

		RE_SIMPLE_RE0.Child = any1
		any1.Next = plus2

		{
			nfaAny := any1.Value.(*Nfa)
			nfaAny.KleenePlus()
			
			RE_SIMPLE_RE0.Value = nfaAny
		}
	case 13:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		plus2 := rhs[1]
		CONCATENATION_RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = any1
		any1.Next = plus2
		plus2.Next = CONCATENATION_RE_SIMPLE_RE3

		{
			nfaAny := any1.Value.(*Nfa)
			nfaAny.KleenePlus()
			rightNfa := CONCATENATION_RE_SIMPLE_RE3.Value.(*Nfa)
			
			nfaAny.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaAny
		}
	case 14:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		plus2 := rhs[1]
		RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = any1
		any1.Next = plus2
		plus2.Next = RE_SIMPLE_RE3

		{
			nfaAny := any1.Value.(*Nfa)
			nfaAny.KleenePlus()
			rightNfa := RE_SIMPLE_RE3.Value.(*Nfa)
			
			nfaAny.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaAny
		}
	case 15:
		RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		star2 := rhs[1]

		RE_SIMPLE_RE0.Child = any1
		any1.Next = star2

		{
			nfaAny := any1.Value.(*Nfa)
			nfaAny.KleeneStar()
			
			RE_SIMPLE_RE0.Value = nfaAny
		}
	case 16:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		star2 := rhs[1]
		CONCATENATION_RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = any1
		any1.Next = star2
		star2.Next = CONCATENATION_RE_SIMPLE_RE3

		{
			nfaAny := any1.Value.(*Nfa)
			nfaAny.KleeneStar()
			rightNfa := CONCATENATION_RE_SIMPLE_RE3.Value.(*Nfa)
			
			nfaAny.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaAny
		}
	case 17:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		any1 := rhs[0]
		star2 := rhs[1]
		RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = any1
		any1.Next = star2
		star2.Next = RE_SIMPLE_RE3

		{
			nfaAny := any1.Value.(*Nfa)
			nfaAny.KleeneStar()
			rightNfa := RE_SIMPLE_RE3.Value.(*Nfa)
			
			nfaAny.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaAny
		}
	case 18:
		RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]

		RE_SIMPLE_RE0.Child = char1

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 19:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]

		CONCATENATION_RE_SIMPLE_RE0.Child = char1
		char1.Next = CONCATENATION_RE_SIMPLE_RE2

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			rightNfa := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
		
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 20:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]

		CONCATENATION_RE_SIMPLE_RE0.Child = char1
		char1.Next = RE_SIMPLE_RE2

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			rightNfa := RE_SIMPLE_RE2.Value.(*Nfa)
		
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 21:
		RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		plus2 := rhs[1]

		RE_SIMPLE_RE0.Child = char1
		char1.Next = plus2

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			newNfa.KleenePlus()
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 22:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		plus2 := rhs[1]
		CONCATENATION_RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = char1
		char1.Next = plus2
		plus2.Next = CONCATENATION_RE_SIMPLE_RE3

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			newNfa.KleenePlus()
			rightNfa := plus2.Value.(*Nfa)
		
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 23:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		plus2 := rhs[1]
		RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = char1
		char1.Next = plus2
		plus2.Next = RE_SIMPLE_RE3

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			newNfa.KleenePlus()
			rightNfa := plus2.Value.(*Nfa)
		
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 24:
		RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		star2 := rhs[1]

		RE_SIMPLE_RE0.Child = char1
		char1.Next = star2

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			newNfa.KleeneStar()
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 25:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		star2 := rhs[1]
		CONCATENATION_RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = char1
		char1.Next = star2
		star2.Next = CONCATENATION_RE_SIMPLE_RE3

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			newNfa.KleeneStar()
			rightNfa := star2.Value.(*Nfa)
		
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 26:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		char1 := rhs[0]
		star2 := rhs[1]
		RE_SIMPLE_RE3 := rhs[2]

		CONCATENATION_RE_SIMPLE_RE0.Child = char1
		char1.Next = star2
		star2.Next = RE_SIMPLE_RE3

		{
			newNfa := newNfaFromChar(char1.Value.(byte))
			newNfa.KleeneStar()
			rightNfa := star2.Value.(*Nfa)
		
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 27:
		SET_ITEMS0 := lhs
		charinset1 := rhs[0]

		SET_ITEMS0.Child = charinset1

		{
			var charSet [256]bool
			charSet[charinset1.Value.(byte)] = true
			
			SET_ITEMS0.Value = charSet
		}
	case 28:
		SET_ITEMS0 := lhs
		charinset1 := rhs[0]
		SET_ITEMS2 := rhs[1]

		SET_ITEMS0.Child = charinset1
		charinset1.Next = SET_ITEMS2

		{
			charSet := SET_ITEMS2.Value.([256]bool)
			charSet[charinset1.Value.(byte)] = true
			
			SET_ITEMS0.Value = charSet
		}
	case 29:
		SET_ITEMS0 := lhs
		charinset1 := rhs[0]
		dash2 := rhs[1]
		charinset3 := rhs[2]

		SET_ITEMS0.Child = charinset1
		charinset1.Next = dash2
		dash2.Next = charinset3

		{
			charStart := charinset1.Value.(byte)
			charEnd := charinset3.Value.(byte)
			
			if charStart > charEnd {
				temp := charStart
				charStart = charEnd
				charEnd = temp
			}
			
			var charSet [256]bool
			for i := charStart; i <= charEnd; i++ {
				charSet[i] = true
			}
			
			SET_ITEMS0.Value = charSet
		}
	case 30:
		SET_ITEMS0 := lhs
		charinset1 := rhs[0]
		dash2 := rhs[1]
		charinset3 := rhs[2]
		SET_ITEMS4 := rhs[3]

		SET_ITEMS0.Child = charinset1
		charinset1.Next = dash2
		dash2.Next = charinset3
		charinset3.Next = SET_ITEMS4

		{
			charStart := charinset1.Value.(byte)
			charEnd := charinset3.Value.(byte)
			charSet := SET_ITEMS4.Value.([256]bool)
			
			if charStart > charEnd {
				temp := charStart
				charStart = charEnd
				charEnd = temp
			}
			
			for i := charStart; i <= charEnd; i++ {
				charSet[i] = true
			}
			
			SET_ITEMS0.Value = charSet
		}
	case 31:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3

		{
			RE_SIMPLE_RE0.Value = CONCATENATION_RE_SIMPLE_RE2.Value
		}
	case 32:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		CONCATENATION_RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = CONCATENATION_RE_SIMPLE_RE4

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			rightNfa := CONCATENATION_RE_SIMPLE_RE4.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 33:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = RE_SIMPLE_RE4

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			rightNfa := RE_SIMPLE_RE4.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 34:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = plus4

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			
			nfaEnclosed.KleenePlus()
			
			RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 35:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = plus4
		plus4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleenePlus()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
				
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 36:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = plus4
		plus4.Next = RE_SIMPLE_RE5

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleenePlus()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
				
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 37:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = star4

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			
			nfaEnclosed.KleeneStar()
			
			RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 38:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = star4
		star4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleeneStar()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 39:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		CONCATENATION_RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = CONCATENATION_RE_SIMPLE_RE2
		CONCATENATION_RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = star4
		star4.Next = RE_SIMPLE_RE5

		{
			nfaEnclosed := CONCATENATION_RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleeneStar()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 40:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3

		{
			RE_SIMPLE_RE0.Value = RE_SIMPLE_RE2.Value
		}
	case 41:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		CONCATENATION_RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = CONCATENATION_RE_SIMPLE_RE4

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			rightNfa := CONCATENATION_RE_SIMPLE_RE4.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 42:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = RE_SIMPLE_RE4

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			rightNfa := RE_SIMPLE_RE4.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 43:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = plus4

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			
			nfaEnclosed.KleenePlus()
			
			RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 44:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = plus4
		plus4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleenePlus()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
				
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 45:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = plus4
		plus4.Next = RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleenePlus()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
				
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 46:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = star4

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			
			nfaEnclosed.KleeneStar()
			
			RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 47:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = star4
		star4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleeneStar()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 48:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_SIMPLE_RE2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_SIMPLE_RE2
		RE_SIMPLE_RE2.Next = rpar3
		rpar3.Next = star4
		star4.Next = RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_SIMPLE_RE2.Value.(*Nfa)
			nfaEnclosed.KleeneStar()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 49:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3

		{
			RE_SIMPLE_RE0.Value = RE_UNION2.Value
		}
	case 50:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		CONCATENATION_RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = CONCATENATION_RE_SIMPLE_RE4

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			rightNfa := CONCATENATION_RE_SIMPLE_RE4.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 51:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = RE_SIMPLE_RE4

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			rightNfa := RE_SIMPLE_RE4.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 52:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = plus4

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			
			nfaEnclosed.KleenePlus()
			
			RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 53:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = plus4
		plus4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			nfaEnclosed.KleenePlus()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
				
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 54:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		plus4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = plus4
		plus4.Next = RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			nfaEnclosed.KleenePlus()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
				
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 55:
		RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]

		RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = star4

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			
			nfaEnclosed.KleeneStar()
			
			RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 56:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = star4
		star4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			nfaEnclosed.KleeneStar()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 57:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		lpar1 := rhs[0]
		RE_UNION2 := rhs[1]
		rpar3 := rhs[2]
		star4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = lpar1
		lpar1.Next = RE_UNION2
		RE_UNION2.Next = rpar3
		rpar3.Next = star4
		star4.Next = RE_SIMPLE_RE5

		{
			nfaEnclosed := RE_UNION2.Value.(*Nfa)
			nfaEnclosed.KleeneStar()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			nfaEnclosed.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = nfaEnclosed
		}
	case 58:
		RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]

		RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 59:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		CONCATENATION_RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = CONCATENATION_RE_SIMPLE_RE4

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			rightNfa := CONCATENATION_RE_SIMPLE_RE4.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 60:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		RE_SIMPLE_RE4 := rhs[3]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = RE_SIMPLE_RE4

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			rightNfa := RE_SIMPLE_RE4.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 61:
		RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		plus4 := rhs[3]

		RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = plus4

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			newNfa.KleenePlus()
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 62:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		plus4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = plus4
		plus4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			newNfa.KleenePlus()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 63:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		plus4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = plus4
		plus4.Next = RE_SIMPLE_RE5

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			newNfa.KleenePlus()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 64:
		RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		star4 := rhs[3]

		RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = star4

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			newNfa.KleeneStar()
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 65:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		star4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = star4
		star4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			newNfa.KleeneStar()
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 66:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		SET_ITEMS2 := rhs[1]
		squarerpar3 := rhs[2]
		star4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = SET_ITEMS2
		SET_ITEMS2.Next = squarerpar3
		squarerpar3.Next = star4
		star4.Next = RE_SIMPLE_RE5

		{
			newNfa := newNfaFromCharClass(SET_ITEMS2.Value.([256]bool))
			newNfa.KleeneStar()
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 67:
		RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]

		RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 68:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		CONCATENATION_RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = CONCATENATION_RE_SIMPLE_RE5

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			rightNfa := CONCATENATION_RE_SIMPLE_RE5.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 69:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		RE_SIMPLE_RE5 := rhs[4]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = RE_SIMPLE_RE5

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			rightNfa := RE_SIMPLE_RE5.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 70:
		RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		plus5 := rhs[4]

		RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = plus5

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			
			newNfa.KleenePlus()
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 71:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		plus5 := rhs[4]
		CONCATENATION_RE_SIMPLE_RE6 := rhs[5]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = plus5
		plus5.Next = CONCATENATION_RE_SIMPLE_RE6

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			newNfa.KleenePlus()
			rightNfa := CONCATENATION_RE_SIMPLE_RE6.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 72:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		plus5 := rhs[4]
		RE_SIMPLE_RE6 := rhs[5]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = plus5
		plus5.Next = RE_SIMPLE_RE6

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			newNfa.KleenePlus()
			rightNfa := RE_SIMPLE_RE6.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 73:
		RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		star5 := rhs[4]

		RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = star5

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			
			newNfa.KleeneStar()
			
			RE_SIMPLE_RE0.Value = &newNfa
		}
	case 74:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		star5 := rhs[4]
		CONCATENATION_RE_SIMPLE_RE6 := rhs[5]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = star5
		star5.Next = CONCATENATION_RE_SIMPLE_RE6

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			newNfa.KleeneStar()
			rightNfa := CONCATENATION_RE_SIMPLE_RE6.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	case 75:
		CONCATENATION_RE_SIMPLE_RE0 := lhs
		squarelpar1 := rhs[0]
		caret2 := rhs[1]
		SET_ITEMS3 := rhs[2]
		squarerpar4 := rhs[3]
		star5 := rhs[4]
		RE_SIMPLE_RE6 := rhs[5]

		CONCATENATION_RE_SIMPLE_RE0.Child = squarelpar1
		squarelpar1.Next = caret2
		caret2.Next = SET_ITEMS3
		SET_ITEMS3.Next = squarerpar4
		squarerpar4.Next = star5
		star5.Next = RE_SIMPLE_RE6

		{
			chars := SET_ITEMS3.Value.([256]bool)
			
			//Skip the first char (empty transition)
			for i := 1; i < len(chars); i++ {
				chars[i] = !chars[i]
			}
			
			newNfa := newNfaFromCharClass(chars)
			newNfa.KleeneStar()
			rightNfa := RE_SIMPLE_RE6.Value.(*Nfa)
			
			newNfa.Concatenate(*rightNfa)
			
			CONCATENATION_RE_SIMPLE_RE0.Value = &newNfa
		}
	}
}
