package main

import "fmt"

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?
func main() {
	numberUnderTest := 600851475143

	//Keep dividing numberUnderTest by incremental i, until it can't be divided by i any longer.
	for i := 2; i <= numberUnderTest; i++ {
		if numberUnderTest%i == 0 && numberUnderTest != i {
			// Divide by factor
			numberUnderTest = numberUnderTest / i
		}
	}
	fmt.Printf("\nThe largest prime factor of 600851475143 is: %d", numberUnderTest)
}
