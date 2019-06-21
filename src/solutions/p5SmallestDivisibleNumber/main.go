package main

import "fmt"

/*
 * https://projecteuler.net/problem=5
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
func main() {
	/*
	 * Since we are given in the question that the smallest divisible
	 * number of 1 to 10 is 2520, and 10 is a factor of 20,
	 * we can start there!
	 */
	nums := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := 2520

	for _, num := range nums {
		result = lcm(result, num)
	}
	fmt.Printf("The smallest positive number that is evenly divided by all of the numbers from 1 to 20 is : %d", result)
}

func gcf(a int, b int) int {
	/*
	 * The Greatest Common Factor (GCF) is the largest positive integer
	 * that divides evenly into the numbers a and b with no remainder.
	 * https://en.wikipedia.org/wiki/Euclidean_algorithm
	 */
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

/*
 * The Lowest Common Multiple (LCM) is defined as:
 * "For two integers a and b, denoted LCM(a,b), the LCM is the smallest
 *  integer that is evenly divisible by both a and b"
 *  https://www.calculatorsoup.com/calculators/math/lcm.php
 *
 * Using the Greatest Common Factor (GCF) approach, the formula is:
*     LCM(a,b) = (a*b)/GCF(a,b)
*
* We could also have used prime factorisation to calculate the LCM,
* but we already used that approach in Problem 3.
*/
func lcm(a int, b int) int {
	return a * b / gcf(a, b)
}
