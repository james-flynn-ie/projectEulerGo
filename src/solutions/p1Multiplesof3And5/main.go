package main

import (
	"fmt"
)

var (
	sum int
)

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	//We can start the loop at 3, seeing as this is the lowest factor.
	for i := 3; i < 1000; i++ {
		//If i is a multiple of 3 or 5, then add it to sum.
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("The sum of all the multiples of 3 or 5 below 1000 is %d", sum)
}
