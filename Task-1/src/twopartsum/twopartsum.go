package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	// find sum of a
	sum := 0
	for _, v := range a {
		sum += v
	}
	// send sum on res
	res <- sum
}

// ConcurrentSum concurrently sums the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	//Get the subtotals from the channel and return their sum
	firstHalf, secondHalf := <-ch, <-ch
	return firstHalf + secondHalf
}

func main() {
	//example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
