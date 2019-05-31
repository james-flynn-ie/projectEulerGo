package main

import "fmt"

//const maxnumber = 600851475143

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?
func main() {
	numberUnderTest := 13195

	//Keep dividing by the lowest prime number, which is 2, until mod 0 is reached.
	for i := 2; i <= numberUnderTest; i++ {
		// A prime factor must be >= 2 & mod = 0
		if numberUnderTest%i == 0 {
			fmt.Println(i)

			// Divide by factor, then reset i to 2 and loop through again.
			numberUnderTest = numberUnderTest / i
			i = 2
		}
	}

}
