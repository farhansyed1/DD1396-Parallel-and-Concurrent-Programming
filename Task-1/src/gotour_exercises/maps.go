package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount finds the word count of a given string
func WordCount(s string) map[string]int {
	// Making a map
	mapOfWords := make(map[string]int)

	// Splitting the string into a slice of substrings
	substrings := strings.Fields(s)

	// Looping through the slice and counting how many times each word appears
	for _, i := range substrings {
		mapOfWords[i]++
	}

	return mapOfWords
}

func main() {
	wc.Test(WordCount)
}
