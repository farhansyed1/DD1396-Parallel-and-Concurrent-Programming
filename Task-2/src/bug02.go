package main

import (
	"fmt"
	"sync"
	"time"
)

// Problem: This program should go to 11, but it seemingly only prints 1 to 10.
// The program has been fixed and now prints 11 as well.
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(11)

	go Print(ch, &wg)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	wg.Wait()
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
		wg.Done()
	}
}
