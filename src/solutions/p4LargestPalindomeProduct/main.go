package main

import (
	"fmt"
	"strconv"
)

var (
	palindromefound   int
	largestpalindrome int
)

//A palindromic number reads the same both ways.
//The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	for i := 999; i >= 0; i-- {
		//j doesn't need to be initialized at 999, because setting it that high would duplicate the previous i values.
		// i * j  is the same as j * i , e.g.: i (999) * j (996) would be the same as i (996) * j (999).
		for j := i; j > 0; j-- {
			numberundertest := i * j

			//Convert to a string before reversing.
			resultstr := strconv.Itoa(numberundertest)
			reversedstr := reverse(resultstr)
			//fmt.Println("\nreversedstr " + reversedstr)

			if resultstr == reversedstr {
				palindromefound = numberundertest
				//Uncomment this line to output the full list of palindromes.
				//fmt.Printf("\nFound a palindrome! %v == %v", resultstr, reversedstr)
			}

			if palindromefound > largestpalindrome {
				largestpalindrome = palindromefound
			}
		}
	}
	fmt.Printf("The largest palindrome made from the product of two 3-digit numbers is: %d", largestpalindrome)
}

func reverse(s string) string {
	//UTF-8 String literals are a sequence of bytes, so we can switch them using this type.
	bytes := []byte(s)

	//Both counters, m and n, start at 0. We increment m upwards through the byte sequence.
	//We decrement n so that it counts from the last [-1] position in the String towards the start.
	//We then swap the byte found at m with the byte at n, and vice versa.
	//Once the counter m becomes equal to n, then we are finished swapping all of the bytes.
	for m, n := 0, len(bytes)-1; m < n; m, n = m+1, n-1 {
		bytes[m], bytes[n] = bytes[n], bytes[m]
	}
	return string(bytes)
}
