package main

import "fmt"

var primefactornumber int

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?
func main() {
	numberUnderTest := 600851475143

	//Keep dividing by the lowest prime number, which is 2, until mod 0 is reached.
	for i := 2; i <= numberUnderTest; i++ {
		if numberUnderTest%i == 0 && numberUnderTest != i {
			// Divide by factor, then reset i to 2 and loop through again.
			numberUnderTest = numberUnderTest / i
			i = 2
		}
	}
	fmt.Printf("\nThe largest prime factor of 600851475143 is: %d", numberUnderTest)
}
