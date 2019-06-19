package main

import "fmt"

var (
	divisiblenumber         int
	smallestdivisiblenumber int
)

/*
 * https://projecteuler.net/problem=5
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := 1

	for _, num := range nums {
		fmt.Println(lcm(result, num))
		result = lcm(result, num)
	}
	fmt.Printf("Final result: %d", result)
}

func gcf(a int, b int) int {
	// Greatest Common Factor https://en.wikipedia.org/wiki/Euclidean_algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcf(a, b)
}
