package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/simoneguidi94/gopapageno/languages/arithmetic"
)

var fname = flag.String("fname", "", "the name of the file to parse")
var numThreads = flag.Int("n", 1, "the number of threads to use")
var numTests = flag.Int("tests", 10, "the number of tests")

func main() {
	//Set the usage message that is printed when incorrect or insufficient arguments are passed
	flag.Usage = func() {
		fmt.Println("Usage: main -fname filename [-n numthreads] [-tests numtests]")
	}

	flag.Parse()

	if *fname == "" || *numThreads < 1 || *numTests < 1 {
		flag.Usage()
		return
	}

	meanAllocTimes := make([]time.Duration, *numThreads)
	meanLexTimes := make([]time.Duration, *numThreads)
	meanParseTimes := make([]time.Duration, *numThreads)

	varianceAllocTimes := make([]time.Duration, *numThreads)
	varianceLexTimes := make([]time.Duration, *numThreads)
	varianceParseTimes := make([]time.Duration, *numThreads)

	stdDevAllocTimes := make([]time.Duration, *numThreads)
	stdDevLexTimes := make([]time.Duration, *numThreads)
	stdDevParseTimes := make([]time.Duration, *numThreads)

	lexPerformanceFile, err := os.Create("lexperformance.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer lexPerformanceFile.Close()

	parsePerformanceFile, err := os.Create("parseperformance.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer parsePerformanceFile.Close()

	lexParallelCodeFile, err := os.Create("lexparallelcode.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer lexParallelCodeFile.Close()
	lexParallelCodeFile.Write([]byte("1, 0\n"))

	parseParallelCodeFile, err := os.Create("parseparallelcode.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer parseParallelCodeFile.Close()
	parseParallelCodeFile.Write([]byte("1, 0\n"))

	for i := 0; i < *numThreads; i++ {
		curNumThreads := i + 1

		memAllocTimes := make([]time.Duration, *numTests)
		lexTimes := make([]time.Duration, *numTests)
		parseTimes := make([]time.Duration, *numTests)

		postfix := ""
		if curNumThreads > 1 {
			postfix = "s"
		}
		fmt.Printf("Parsing file %s (%d tests, %d thread%s)\n\n", path.Base(*fname), *numTests, curNumThreads, postfix)
		for j := 0; j < *numTests; j++ {
			fmt.Printf("Test n° %d:\n", j+1)

			root, err := arithmetic.ParseFile(*fname, curNumThreads)

			runtime.GC()

			if err == nil {
				fmt.Println("Parse succeded!")
				fmt.Printf("Time to alloc memory: %s\n", arithmetic.Stats.AllocMemTime)
				fmt.Printf("Time to lex: %s\n", arithmetic.Stats.LexTimeTotal)
				fmt.Printf("Time to parse: %s\n\n", arithmetic.Stats.ParseTimeTotal)
				fmt.Printf("Result: %d\n", *root.Value.(*int64))

				memAllocTimes[j] = arithmetic.Stats.AllocMemTime
				lexTimes[j] = arithmetic.Stats.LexTimeTotal
				parseTimes[j] = arithmetic.Stats.ParseTimeTotal
			} else {
				//This should not happen
				fmt.Println("Parse failed!")
				fmt.Println(err.Error())
				return
			}
		}

		meanAllocTimes[i] = calculateMean(memAllocTimes)
		meanLexTimes[i] = calculateMean(lexTimes)
		meanParseTimes[i] = calculateMean(parseTimes)

		fmt.Printf("Mean time to alloc memory: %s\n", meanAllocTimes[i])
		fmt.Printf("Mean time to lex: %s\n", meanLexTimes[i])
		fmt.Printf("Mean time to parse: %s\n\n", meanParseTimes[i])

		varianceAllocTimes[i] = calculateVariance(memAllocTimes)
		varianceLexTimes[i] = calculateVariance(lexTimes)
		varianceParseTimes[i] = calculateVariance(parseTimes)

		fmt.Printf("Variance in memory allocation: %s\n", varianceAllocTimes[i])
		fmt.Printf("Variance in lexing: %s\n", varianceLexTimes[i])
		fmt.Printf("Variance in parsing: %s\n\n", varianceParseTimes[i])

		stdDevAllocTimes[i] = calculateStdDev(memAllocTimes)
		stdDevLexTimes[i] = calculateStdDev(lexTimes)
		stdDevParseTimes[i] = calculateStdDev(parseTimes)

		fmt.Printf("Standard deviation in memory allocation: %s\n", stdDevAllocTimes[i])
		fmt.Printf("Standard deviation in lexing: %s\n", stdDevLexTimes[i])
		fmt.Printf("Standard deviation in parsing: %s\n\n", stdDevParseTimes[i])

		lexPerformanceFile.Write([]byte(fmt.Sprintf("%d, %.2f, %.2f\n", curNumThreads, meanLexTimes[i].Seconds()*1000, stdDevLexTimes[i].Seconds()*1000)))
		parsePerformanceFile.Write([]byte(fmt.Sprintf("%d, %.2f, %.2f\n", curNumThreads, meanParseTimes[i].Seconds()*1000, stdDevParseTimes[i].Seconds()*1000)))

		if curNumThreads > 1 {
			parallelLex := 1 - (1/(meanLexTimes[0].Seconds()/meanLexTimes[i].Seconds())-1/float64(curNumThreads))/(1-1/float64(curNumThreads))
			parallelParse := 1 - (1/(meanParseTimes[0].Seconds()/meanParseTimes[i].Seconds())-1/float64(curNumThreads))/(1-1/float64(curNumThreads))

			fmt.Printf("Parallel percent in lexing: %.1f%%\n", parallelLex*100)
			fmt.Printf("Parallel percent in parsing: %.1f%%\n", parallelParse*100)

			lexParallelCodeFile.Write([]byte(fmt.Sprintf("%d, %.1f\n", curNumThreads, parallelLex*100)))
			parseParallelCodeFile.Write([]byte(fmt.Sprintf("%d, %.1f\n", curNumThreads, parallelParse*100)))
		}

	}
}

func calculateMean(durations []time.Duration) time.Duration {
	total := time.Duration(0)

	for i := 0; i < len(durations); i++ {
		total += durations[i]
	}

	return total / time.Duration(len(durations))
}

func calculateVariance(durations []time.Duration) time.Duration {
	n := float64(0)
	mean := float64(0)
	M2 := float64(0)

	for _, d := range durations {
		n += 1
		x := d.Seconds()
		delta := x - mean
		mean += delta / n
		delta2 := x - mean
		M2 += delta * delta2
	}

	if n < 2 {
		return time.Duration(0)
	} else {
		return time.Duration((M2 / (n - 1)) * 1e9)
	}
}

func calculateStdDev(durations []time.Duration) time.Duration {
	n := float64(0)
	mean := float64(0)
	M2 := float64(0)

	for _, d := range durations {
		n += 1
		x := d.Seconds()
		delta := x - mean
		mean += delta / n
		delta2 := x - mean
		M2 += delta * delta2
	}

	if n < 2 {
		return time.Duration(0)
	} else {
		return time.Duration(math.Sqrt(M2/(n-1)) * 1e9)
	}
}
