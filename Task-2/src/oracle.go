// Stefan Nilsson 2013-03-13
// Farhan Syed 2022-03-30

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	// Answer questions asked by user.
	go func() {
		for {
			go prophecy(<-questions, answers)
		}
	}()

	// Make random prophecies
	go func() {
		for {
			time.Sleep(time.Duration(15) * time.Second)
			prophecy("", answers)
		}
	}()

	// Print answers to questions and print prophecies
	go func() {
		for {
			printAnswers(answers)
		}
	}()

	return questions
}

// Prints the Oracle's answers with a short delay
func printAnswers(answers <-chan string) {
	//Loop through each answer
	for answer := range answers {
		//Loop through each character
		for _, char := range answer {
			fmt.Print(string(char))
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"The earth is flat.",
		"The galaxy is very small.",
	}
	for _, w := range words {
		if w == "Hello" {
			answer <- "Wassup bro"
		} else if w == "Water" {
			answer <- "Water is wet"
		} else if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	if longestWord == "" {
		answer <- nonsense[rand.Intn(len(nonsense))]
	} else {
		answer <- longestWord + ". How wonderful! " + nonsense[rand.Intn(len(nonsense))]
	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
