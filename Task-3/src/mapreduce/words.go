// Farhan Syed 2022-04-07
// This program finds the word count of a text file using parallel go routines
package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
//
func WordCount(text string) map[string]int {
	// Editing text
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, ",", " ")

	results := make(chan map[string]int) // Channel used to communicate between go routines
	freqs := make(map[string]int)
	workers := 15 // The number of go routines

	//Map function - splitting string into smaller substrings
	sliceOfWords := strings.Fields(text)
	sizeOfOneSubstring := len(sliceOfWords) / workers
	leftBound, rightBound := 0, sizeOfOneSubstring

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		if i == workers-1 { // The last section gets all remaining words
			rightBound = len(sliceOfWords)
			leftBound += sizeOfOneSubstring
		} else if i != 0 { // Moving boundaries to next selection of words
			rightBound += sizeOfOneSubstring
			leftBound += sizeOfOneSubstring
		}
		go func(leftBound, rightBound int) {
			stringForRoutine := sliceOfWords[leftBound:rightBound]
			mapForRoutine := make(map[string]int)
			for _, word := range stringForRoutine {
				mapForRoutine[word]++
			}
			results <- mapForRoutine
			wg.Done()
		}(leftBound, rightBound)
	}

	//Reduce function - adding the word counts of each map to freqs
	for i := 0; i < workers; i++ {
		for word, count := range <-results {
			freqs[word] += count
		}
	}
	wg.Wait()
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	data, _ := ioutil.ReadFile(DataFile)
	text := string(data)
	fmt.Printf("%#v", WordCount(text))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
