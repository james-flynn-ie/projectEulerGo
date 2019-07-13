package main

import "fmt"

var highestprimenumber int

const maxcountervalue = 10001

/*
 * https://projecteuler.net/problem=7
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 * What is the 10 001st prime number?
 */
func main() {
	// For a description on how to use channels, see: https://tour.golang.org/concurrency/2
	ch := make(chan int)
	go generatenumbersequence(ch)

	// Start at 2, as we are told in the problem that 1 isn't a prime.
	for i := 2; i <= maxcountervalue; i++ {
		numberundertest := <-ch
		ch1 := make(chan int)

		// Take the next number in the channel, and check if it's a prime.
		go filterprimenumbers(ch, ch1, numberundertest)
		ch = ch1
	}

	highestprimenumber := <-ch
	fmt.Printf("\nThe %dth prime factor is: %d", maxcountervalue, highestprimenumber)
}

func filterprimenumbers(in <-chan int, out chan<- int, numberundertest int) {
	/*
	 * Keep checking the next int in the channel to see if its a prime.
	 * If not, discard it by redirecting to out channel.
	 */
	for {
		i := <-in
		if i%numberundertest != 0 {
			out <- i
		}
	}
}

func generatenumbersequence(ch chan<- int) {
	// Write incrementing ints to the channel, for checking if the next number is prime.
	for i := 2; ; i++ {
		ch <- i
	}
}
