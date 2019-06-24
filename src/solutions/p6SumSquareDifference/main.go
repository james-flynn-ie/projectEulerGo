package main

import "fmt"

var (
	difference   int
	squaredsum   int
	sumofnums    int
	sumofsquares int
)

const maxnumber = 100

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
		/*
		* Find the sum of all the squares from 1 to the highest natural number we
		* are seeking in the sumofsquares var.
		 */
		sumofsquares += i * i

		/*
		* We also want to keep a sum of all the numbers (unsquared) for the next
		* step, where we square the sum of all the numbers from 1 to the
		* highest natural number we are seeking.
		 */
		sumofnums += i
	}
	fmt.Printf("The sum of the squares of the first %v natural numbers is: %v", maxnumber, sumofsquares)

	squaredsum = sumofnums * sumofnums
	fmt.Printf("\nThe squared sum of the squares of the first %v natural numbers is: %v", maxnumber, squaredsum)

	difference = squaredsum - sumofsquares
	fmt.Printf("\ndifference between the sum of the squares and the square of the sum for the first %v natural numbers is: %v", maxnumber, difference)
}
