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

	success, _ := xml.ParseFile(*fname, *numThreads)

	if success {
		fmt.Println("Parse succeded!")
		fmt.Printf("Stack pool size: %d\n", xml.Stats.StackPoolSize)
		fmt.Printf("StackPtr pool size: %d\n", xml.Stats.StackPtrPoolSize)
		fmt.Printf("Time to alloc memory: %s\n", xml.Stats.AllocMemTime)
		fmt.Printf("Time to lex: %s\n", xml.Stats.LexTime)
		fmt.Printf("Number of tokens: %d\n", xml.Stats.NumTokens)
		fmt.Printf("Time to parse: %s\n", xml.Stats.ParseTime)
		fmt.Printf("Remaining parser stacks: %d\n", xml.Stats.RemainingStacks)
		fmt.Printf("Remaining parser stackptrs: %d\n", xml.Stats.RemainingStackPtrs)

	} else {
		fmt.Println("Parse failed!")
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
