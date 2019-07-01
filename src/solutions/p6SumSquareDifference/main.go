package main

import "fmt"

var (
	difference   int
	squaredsum   int
	sumofnums    int
	sumofsquares int
)

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
	const N = 100

	sumofnums = calculatesumofnums(N)
	squaredsum = sumofnums * sumofnums
	fmt.Printf("The squared sum of the squares of the first %v natural numbers is: %v", N, squaredsum)

	sumofsquares = calculatesumofsquare(N)
	fmt.Printf("\nThe sum of the squares of the first %v natural numbers is: %v", N, sumofsquares)

	difference = squaredsum - sumofsquares
	fmt.Printf("\ndifference between the sum of the squares and the square of the sum for the first %v natural numbers is: %v", N, difference)
}

func calculatesumofnums(N int) int {
	/*
	 * We can use the following Sum of Natural Numbers algorithm:
	 * https: //trans4mind.com/personal_development/mathematics/series/sumNaturalNumbers.htm#mozTocId914933
	 */
	sumofnums = N * (N + 1) / 2
	return sumofnums
}

func calculatesumofsquare(N int) int {
	/*
	 * We can use the following Sum of the Squares of Natural Numbers algorithm:
	 * https://trans4mind.com/personal_development/mathematics/series/sumNaturalSquares.htm
	 */
	sumofsquares = (N * (N + 1) * (2*N + 1)) / 6
	return sumofsquares
}
