package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/simoneguidi94/gopapageno/languages/xml"
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
		xml.SetCPUProfileFile(cpuprofileFile)
	}

	fmt.Println("Available cores:", runtime.GOMAXPROCS(0))

	fmt.Println("Number of threads:", *numThreads)

	_, err := xml.ParseFile(*fname, *numThreads)

	if err == nil {
		fmt.Println("Parse succeded!")
		for i, v := range xml.Stats.StackPoolSizes {
			fmt.Printf("Stack pool size (thread %d): %d\n", i, v)
		}
		for i, v := range xml.Stats.StackPoolNewNonterminalsSizes {
			fmt.Printf("Stack pool new nonterminals size (thread %d): %d\n", i, v)
		}
		for i, v := range xml.Stats.StackPtrPoolSizes {
			fmt.Printf("StackPtr pool size (thread %d): %d\n", i, v)
		}
		fmt.Printf("Stack pool final pass size: %d\n", xml.Stats.StackPoolSizeFinalPass)
		fmt.Printf("Stack pool final pass new nonterminals size: %d\n", xml.Stats.StackPoolNewNonterminalsSizeFinalPass)
		fmt.Printf("StackPtr pool final pass size: %d\n", xml.Stats.StackPtrPoolSizeFinalPass)
		fmt.Printf("Time to alloc memory: %s\n\n", xml.Stats.AllocMemTime)

		for i, v := range xml.Stats.CutPoints {
			fmt.Printf("cutpoint %d: %d\n", i, v)
		}
		for i, v := range xml.Stats.LexTimes {
			fmt.Printf("Time to lex (thread %d): %s\n", i, v)
		}
		fmt.Printf("Time to lex (total): %s\n\n", xml.Stats.LexTimeTotal)

		for i, v := range xml.Stats.NumTokens {
			fmt.Printf("Number of tokens (thread %d): %d\n", i, v)
		}
		fmt.Printf("Number of tokens (total): %d\n", xml.Stats.NumTokensTotal)
		for i, v := range xml.Stats.ParseTimes {
			fmt.Printf("Time to parse (thread %d): %s\n", i, v)
		}
		fmt.Printf("Time to recombine the stacks: %s\n", xml.Stats.RecombiningStacksTime)
		fmt.Printf("Time to parse (final pass): %s\n", xml.Stats.ParseTimeFinalPass)
		fmt.Printf("Time to parse (total): %s\n\n", xml.Stats.ParseTimeTotal)

		for i, v := range xml.Stats.RemainingStacks {
			fmt.Printf("Remaining stacks (thread %d): %d\n", i, v)
		}
		for i, v := range xml.Stats.RemainingStacksNewNonterminals {
			fmt.Printf("Remaining stacks new nonterminals (thread %d): %d\n", i, v)
		}
		for i, v := range xml.Stats.RemainingStackPtrs {
			fmt.Printf("Remaining stackPtrs (thread %d): %d\n", i, v)
		}

		fmt.Printf("Remaining stacks final pass: %d\n", xml.Stats.RemainingStacksFinalPass)
		fmt.Printf("Remaining stacks new nonterminals final pass: %d\n", xml.Stats.RemainingStacksNewNonterminalsFinalPass)
		fmt.Printf("Remaining stackPtrs final pass: %d\n", xml.Stats.RemainingStackPtrsFinalPass)

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
