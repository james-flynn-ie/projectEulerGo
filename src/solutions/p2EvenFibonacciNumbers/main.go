package main

import "fmt"

var (
	lastfibonaccinumber    int
	currentfibonaccinumber int
	nextfibonaccinumber    int
	sum                    int
)

//Each new term in the Fibonacci sequence is generated by adding the previous two terms.
//By starting with 1 and 2, the first 10 terms will be:
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//By considering the terms in the Fibonacci sequence whose values do not exceed four million,
//find the sum of the even-valued terms.
func main() {
	//Initialize fibonacci number and sum vars with the values we are given in the question
	lastfibonaccinumber = 1
	currentfibonaccinumber = 2

	// We must'nt forget to add the first even term to the sum!
	sum += currentfibonaccinumber

	for nextfibonaccinumber < 4000000 {
		nextfibonaccinumber = lastfibonaccinumber + currentfibonaccinumber

		//Only add the even values to the sum.
		if nextfibonaccinumber%2 == 0 {
			sum += nextfibonaccinumber
		}

		// Reassign the fibonacci numbers for the next loop.
		lastfibonaccinumber = currentfibonaccinumber
		currentfibonaccinumber = nextfibonaccinumber
	}
	fmt.Printf("The total sum of all even Fibonacci numbers under four million is : %d", sum)
}
