package main

import "fmt"

var highestprimenumber int

const maxcountervalue = 6

/*
 * https://projecteuler.net/problem=7
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 * What is the 10 001st prime number?
 */
func main() {
	// For a description on how to use channels, see: https://tour.golang.org/concurrency/2
	ch := make(chan int)
	go generatenumbersequence(ch)

	// Start at 2, as 1 is known to be a prime.
	for i := 2; i <= maxcountervalue; i++ {
		prime := <-ch
		ch1 := make(chan int)
		go filterprimenumbers(ch, ch1, prime)
		ch = ch1
	}

	highestprimenumber := <-ch
	fmt.Printf("\nThe %dth prime factor is: %d", maxcountervalue, highestprimenumber)
}

func filterprimenumbers(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func generatenumbersequence(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}
