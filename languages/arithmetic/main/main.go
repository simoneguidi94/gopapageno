package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/simoneguidi94/gopapageno/languages/arithmetic"
)

var cpuprofile = flag.String("cpuprofile", "", "") //write cpu profile to file")
var memprofile = flag.String("memprofile", "", "") //write memory profile to file")

var cpuprofileFile *os.File

var fname = flag.String("fname", "", "the name of the file to parse")
var numThreads = flag.Int("n", 1, "the number of threads to use")

func main() {
	//Set flags (for debugging only)
	//flag.Set("fname", "languages/arithmetic/data/small.txt")
	//flag.Set("n", "2")

	//Set the usage message that is printed when incorrect or insufficient arguments are passed
	flag.Usage = func() {
		fmt.Println("Usage: main -fname filename [-n numthreads]")
	}

	flag.Parse()

	if *fname == "" || *numThreads < 1 {
		flag.Usage()
		return
	}

	//Code needed for the cpu profiler
	if *cpuprofile != "" {
		err := error(nil)
		cpuprofileFile, err = os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		arithmetic.SetCPUProfileFile(cpuprofileFile)
	}

	fmt.Println("Available cores:", runtime.GOMAXPROCS(0))

	fmt.Println("Number of threads:", *numThreads)

	root, err := arithmetic.ParseFile(*fname, *numThreads)

	if err == nil {
		fmt.Println("Parse succeded!")
		for i, v := range arithmetic.Stats.StackPoolSizes {
			fmt.Printf("Stack pool size (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic.Stats.StackPoolNewNonterminalsSizes {
			fmt.Printf("Stack pool new nonterminals size (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic.Stats.StackPtrPoolSizes {
			fmt.Printf("StackPtr pool size (thread %d): %d\n", i, v)
		}
		fmt.Printf("Stack pool final pass size: %d\n", arithmetic.Stats.StackPoolSizeFinalPass)
		fmt.Printf("Stack pool final pass new nonterminals size: %d\n", arithmetic.Stats.StackPoolNewNonterminalsSizeFinalPass)
		fmt.Printf("StackPtr pool final pass size: %d\n", arithmetic.Stats.StackPtrPoolSizeFinalPass)
		fmt.Printf("Time to alloc memory: %s\n\n", arithmetic.Stats.AllocMemTime)

		for i, v := range arithmetic.Stats.CutPoints {
			fmt.Printf("cutpoint %d: %d\n", i, v)
		}
		for i, v := range arithmetic.Stats.LexTimes {
			fmt.Printf("Time to lex (thread %d): %s\n", i, v)
		}
		fmt.Printf("Time to lex (total): %s\n\n", arithmetic.Stats.LexTimeTotal)

		for i, v := range arithmetic.Stats.NumTokens {
			fmt.Printf("Number of tokens (thread %d): %d\n", i, v)
		}
		fmt.Printf("Number of tokens (total): %d\n", arithmetic.Stats.NumTokensTotal)
		for i, v := range arithmetic.Stats.ParseTimes {
			fmt.Printf("Time to parse (thread %d): %s\n", i, v)
		}
		fmt.Printf("Time to recombine the stacks: %s\n", arithmetic.Stats.RecombiningStacksTime)
		fmt.Printf("Time to parse (final pass): %s\n", arithmetic.Stats.ParseTimeFinalPass)
		fmt.Printf("Time to parse (total): %s\n\n", arithmetic.Stats.ParseTimeTotal)

		for i, v := range arithmetic.Stats.RemainingStacks {
			fmt.Printf("Remaining stacks (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic.Stats.RemainingStacksNewNonterminals {
			fmt.Printf("Remaining stacks new nonterminals (thread %d): %d\n", i, v)
		}
		for i, v := range arithmetic.Stats.RemainingStackPtrs {
			fmt.Printf("Remaining stackPtrs (thread %d): %d\n", i, v)
		}

		fmt.Printf("Remaining stacks final pass: %d\n", arithmetic.Stats.RemainingStacksFinalPass)
		fmt.Printf("Remaining stacks new nonterminals final pass: %d\n", arithmetic.Stats.RemainingStacksNewNonterminalsFinalPass)
		fmt.Printf("Remaining stackPtrs final pass: %d\n\n", arithmetic.Stats.RemainingStackPtrsFinalPass)

		fmt.Printf("Result: %d\n", *root.Value.(*int64))
	} else {
		fmt.Println("Parse failed!")
		fmt.Println(err.Error())
	}

	//Code needed for the mem profiler
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		//runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
