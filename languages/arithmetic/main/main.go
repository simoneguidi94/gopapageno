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
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Available cores:", runtime.GOMAXPROCS(0))

	fmt.Println("Number of threads:", *numThreads)

	success, result := arithmetic.Parse(*fname, *numThreads)

	if success {
		fmt.Println("Parse succeded!")
		fmt.Println("Result:", result)
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
