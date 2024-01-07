package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
// The program is now fixed by using another go routine
func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}
