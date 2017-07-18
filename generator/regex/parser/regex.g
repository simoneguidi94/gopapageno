/*
parserPreallocMem initializes all the memory pools required by the semantic function.
*/
func parserPreallocMem(numThreads int) {
}
%%

%axiom RE

%%

RE : UNION
{
	$$.Value = $1.Value
} | SIMPLE_RE
{
	$$.Value = $1.Value
};
UNION : RE pipe SIMPLE_RE
{
	leftNfa := $1.Value.(*Nfa)
	rightNfa := $3.Value.(*Nfa)
	
	leftNfa.Unite(*rightNfa)
	
	$$.Value = leftNfa
};
SIMPLE_RE : CONCATENATION
{
	$$.Value = $1.Value
} | any
{
	$$.Value = $1.Value
} | char
{
	newNfa := newNfaFromChar($1.Value.(byte))
	
	$$.Value = &newNfa
} | squarelpar SET_ITEMS squarerpar
{
	newNfa := newNfaFromCharClass($2.Value.([256]bool))
	
	$$.Value = &newNfa
} | squarelpar caret SET_ITEMS squarerpar
{
	chars := $3.Value.([256]bool)
	
	//Skip the first char (empty transition)
	for i := 1; i < len(chars); i++ {
		chars[i] = !chars[i]
	}
	
	newNfa := newNfaFromCharClass(chars)
	
	$$.Value = &newNfa
} | lpar RE rpar
{
	$$.Value = $2.Value
} | any star
{
	nfaAny := $1.Value.(*Nfa)
	nfaAny.KleeneStar()
	
	$$.Value = nfaAny
} | char star
{
	newNfa := newNfaFromChar($1.Value.(byte))
	newNfa.KleeneStar()
	
	$$.Value = &newNfa
} | squarelpar SET_ITEMS squarerpar star
{
	newNfa := newNfaFromCharClass($2.Value.([256]bool))
	newNfa.KleeneStar()
	
	$$.Value = &newNfa
} | squarelpar caret SET_ITEMS squarerpar star
{
	chars := $3.Value.([256]bool)
	
	//Skip the first char (empty transition)
	for i := 1; i < len(chars); i++ {
		chars[i] = !chars[i]
	}
	
	newNfa := newNfaFromCharClass(chars)
	
	newNfa.KleeneStar()
	
	$$.Value = &newNfa
} | lpar RE rpar star
{
	nfaEnclosed := $2.Value.(*Nfa)
	
	nfaEnclosed.KleeneStar()
	
	$$.Value = nfaEnclosed
} | any plus
{
	nfaAny := $1.Value.(*Nfa)
	nfaAny.KleenePlus()
	
	$$.Value = nfaAny
} | char plus
{
	newNfa := newNfaFromChar($1.Value.(byte))
	newNfa.KleenePlus()
	
	$$.Value = &newNfa
} | squarelpar SET_ITEMS squarerpar plus
{
	newNfa := newNfaFromCharClass($2.Value.([256]bool))
	newNfa.KleenePlus()
	
	$$.Value = &newNfa
} | squarelpar caret SET_ITEMS squarerpar plus
{
	chars := $3.Value.([256]bool)
	
	//Skip the first char (empty transition)
	for i := 1; i < len(chars); i++ {
		chars[i] = !chars[i]
	}
	
	newNfa := newNfaFromCharClass(chars)
	
	newNfa.KleenePlus()
	
	$$.Value = &newNfa
} | lpar RE rpar plus
{
	nfaEnclosed := $2.Value.(*Nfa)
	
	nfaEnclosed.KleenePlus()
	
	$$.Value = nfaEnclosed
};


CONCATENATION : any SIMPLE_RE
{
	leftNfa := $1.Value.(*Nfa)
	rightNfa := $2.Value.(*Nfa)
	
	leftNfa.Concatenate(*rightNfa)
	
	$$.Value = leftNfa
} | char SIMPLE_RE
{
	newNfa := newNfaFromChar($1.Value.(byte))
	rightNfa := $2.Value.(*Nfa)

	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | squarelpar SET_ITEMS squarerpar SIMPLE_RE
{
	newNfa := newNfaFromCharClass($2.Value.([256]bool))
	rightNfa := $4.Value.(*Nfa)
	
	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | squarelpar caret SET_ITEMS squarerpar SIMPLE_RE
{
	chars := $3.Value.([256]bool)
	
	//Skip the first char (empty transition)
	for i := 1; i < len(chars); i++ {
		chars[i] = !chars[i]
	}
	
	newNfa := newNfaFromCharClass(chars)
	rightNfa := $5.Value.(*Nfa)
	
	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | lpar RE rpar SIMPLE_RE
{
	nfaEnclosed := $2.Value.(*Nfa)
	rightNfa := $4.Value.(*Nfa)
	
	nfaEnclosed.Concatenate(*rightNfa)
	
	$$.Value = nfaEnclosed
} | any star SIMPLE_RE
{
	nfaAny := $1.Value.(*Nfa)
	nfaAny.KleeneStar()
	rightNfa := $3.Value.(*Nfa)
	
	nfaAny.Concatenate(*rightNfa)
	
	$$.Value = nfaAny
} | char star SIMPLE_RE
{
	newNfa := newNfaFromChar($1.Value.(byte))
	newNfa.KleeneStar()
	rightNfa := $2.Value.(*Nfa)

	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | squarelpar SET_ITEMS squarerpar star SIMPLE_RE
{
	newNfa := newNfaFromCharClass($2.Value.([256]bool))
	newNfa.KleeneStar()
	rightNfa := $5.Value.(*Nfa)
	
	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | squarelpar caret SET_ITEMS squarerpar star SIMPLE_RE
{
	chars := $3.Value.([256]bool)
	
	//Skip the first char (empty transition)
	for i := 1; i < len(chars); i++ {
		chars[i] = !chars[i]
	}
	
	newNfa := newNfaFromCharClass(chars)
	newNfa.KleeneStar()
	rightNfa := $6.Value.(*Nfa)
	
	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | lpar RE rpar star SIMPLE_RE
{
	nfaEnclosed := $2.Value.(*Nfa)
	nfaEnclosed.KleeneStar()
	rightNfa := $5.Value.(*Nfa)
	
	nfaEnclosed.Concatenate(*rightNfa)
	
	$$.Value = nfaEnclosed
} | any plus SIMPLE_RE
{
	nfaAny := $1.Value.(*Nfa)
	nfaAny.KleenePlus()
	rightNfa := $3.Value.(*Nfa)
	
	nfaAny.Concatenate(*rightNfa)
	
	$$.Value = nfaAny
} | char plus SIMPLE_RE
{
	newNfa := newNfaFromChar($1.Value.(byte))
	newNfa.KleenePlus()
	rightNfa := $2.Value.(*Nfa)

	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | squarelpar SET_ITEMS squarerpar plus SIMPLE_RE
{
	newNfa := newNfaFromCharClass($2.Value.([256]bool))
	newNfa.KleenePlus()
	rightNfa := $5.Value.(*Nfa)
	
	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | squarelpar caret SET_ITEMS squarerpar plus SIMPLE_RE
{
	chars := $3.Value.([256]bool)
	
	//Skip the first char (empty transition)
	for i := 1; i < len(chars); i++ {
		chars[i] = !chars[i]
	}
	
	newNfa := newNfaFromCharClass(chars)
	newNfa.KleenePlus()
	rightNfa := $6.Value.(*Nfa)
	
	newNfa.Concatenate(*rightNfa)
	
	$$.Value = &newNfa
} | lpar RE rpar plus SIMPLE_RE
{
	nfaEnclosed := $2.Value.(*Nfa)
	nfaEnclosed.KleenePlus()
	rightNfa := $5.Value.(*Nfa)
	
	nfaEnclosed.Concatenate(*rightNfa)
		
	$$.Value = nfaEnclosed
};
SET_ITEMS : charinset
{
	var charSet [256]bool
	charSet[$1.Value.(byte)] = true
	
	$$.Value = charSet
} | charinset dash charinset
{
	charStart := $1.Value.(byte)
	charEnd := $3.Value.(byte)
	
	if charStart > charEnd {
		temp := charStart
		charStart = charEnd
		charEnd = temp
	}
	
	var charSet [256]bool
	for i := charStart; i <= charEnd; i++ {
		charSet[i] = true
	}
	
	$$.Value = charSet
} | charinset SET_ITEMS
{
	charSet := $2.Value.([256]bool)
	charSet[$1.Value.(byte)] = true
	
	$$.Value = charSet
} | charinset dash charinset SET_ITEMS
{
	charStart := $1.Value.(byte)
	charEnd := $3.Value.(byte)
	charSet := $4.Value.([256]bool)
	
	if charStart > charEnd {
		temp := charStart
		charStart = charEnd
		charEnd = temp
	}
	
	for i := charStart; i <= charEnd; i++ {
		charSet[i] = true
	}
	
	$$.Value = charSet
};