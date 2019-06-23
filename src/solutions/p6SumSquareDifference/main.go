package main

import "fmt"

var (
	squaredsum   int
	sumofnums    int
	sumofsquares int
)

const maxnumber = 10

/*
 * https://projecteuler.net/problem=6
 * The sum of the squares of the first ten natural numbers is,
 *       1^2 + 2^2 + ... + 10^2 = 385
 * The square of the sum of the first ten natural numbers is,
 *      (1 + 2 + ... + 10)^2 = 55^2 = 3025
 * Hence the difference between the sum of the squares of the first ten natural
 *  numbers and the square of the sum is 3025 − 385 = 2640.
 * Find the difference between the sum of the squares of the first one hundred
 * natural numbers and the square of the sum.
 */
func main() {
	for i := 1; i <= maxnumber; i++ {
		sumofnums += i
		sumofsquares += i * i
	}
	fmt.Printf("The sum of the squares of the first %v natural numbers is: %v", maxnumber, sumofsquares)

	squaredsum = sumofnums * sumofnums
	fmt.Printf("\nThe squared sum of the first %v natural numbers is: %v", maxnumber, squaredsum)
}
