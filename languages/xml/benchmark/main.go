package main

import (
	"flag"
	"fmt"
	"math"
	"path"
	"time"

	"github.com/simoneguidi94/gopapageno/languages/xml"
)

var fname = flag.String("fname", "", "the name of the file to parse")
var numThreads = flag.Int("n", 1, "the number of threads to use")

func main() {
	//Set the usage message that is printed when incorrect or insufficient arguments are passed
	flag.Usage = func() {
		fmt.Println("Usage: main -fname filename [-n numthreads]")
	}

	flag.Parse()

	if *fname == "" || *numThreads < 1 {
		flag.Usage()
		return
	}

	const NUM_TESTS int = 10

	memAllocTimes := make([]time.Duration, NUM_TESTS)
	lexTimes := make([]time.Duration, NUM_TESTS)
	parseTimes := make([]time.Duration, NUM_TESTS)

	postfix := ""
	if *numThreads > 1 {
		postfix = "s"
	}
	fmt.Printf("Parsing file %s (%d tests, %d thread%s)\n\n", path.Base(*fname), NUM_TESTS, *numThreads, postfix)
	for i := 0; i < NUM_TESTS; i++ {
		fmt.Printf("Test n° %d:\n", i+1)

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
			fmt.Printf("Remaining parser stackptrs: %d\n\n", xml.Stats.RemainingStackPtrs)

			memAllocTimes[i] = xml.Stats.AllocMemTime
			lexTimes[i] = xml.Stats.LexTime
			parseTimes[i] = xml.Stats.ParseTime
		} else {
			//This should not happen
			fmt.Println("Parse failed!")
			return
		}
	}

	meanAllocTime := calculateMean(memAllocTimes)
	meanLexTime := calculateMean(lexTimes)
	meanParseTime := calculateMean(parseTimes)

	fmt.Printf("Mean time to alloc memory: %s\n", meanAllocTime)
	fmt.Printf("Mean time to lex: %s\n", meanLexTime)
	fmt.Printf("Mean time to parse: %s\n\n", meanParseTime)

	stdDevAllocTime := calculateStandardDeviation(memAllocTimes)
	stdDevLexTime := calculateStandardDeviation(lexTimes)
	stdDevParseTime := calculateStandardDeviation(parseTimes)

	fmt.Printf("Standard deviation in memory allocation: %s\n", stdDevAllocTime)
	fmt.Printf("Standard deviation in lexing: %s\n", stdDevLexTime)
	fmt.Printf("Standard deviation in parsing: %s\n\n", stdDevParseTime)
}

func calculateMean(durations []time.Duration) time.Duration {
	total := time.Duration(0)

	for i := 0; i < len(durations); i++ {
		total += durations[i]
	}

	return total / time.Duration(len(durations))
}

func calculateStandardDeviation(durations []time.Duration) time.Duration {
	total := float64(0)
	mean := calculateMean(durations).Seconds()

	for i := 0; i < len(durations); i++ {
		total += (durations[i].Seconds() - mean) * (durations[i].Seconds() - mean)
	}

	return time.Duration((math.Sqrt(total) / float64(len(durations))) * 1e9)
}
