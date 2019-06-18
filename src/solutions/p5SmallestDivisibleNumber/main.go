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
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for _, num := range nums {
		fmt.Println(lcm(20, num))
	}
}

func gcd(a int, b int) int {
	// Greatest Common Divisor https://en.wikipedia.org/wiki/Euclidean_algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

// Python:
// def gcd(a, b):
//     while b != 0:
//         a, b = b, a % b
//     return a

// def lcm(a, b):
//     return a*b/gcd(a, b)

// print reduce(lcm, range(1, 20+1))


// Java
// public class Euler5 {
//     final static int MAX_N = 20;

//     static int gcd(int a, int b) {
//         while (b != 0) {
//             int t = a;
//             a = b;
//             b = t % b;
//         }
//         return a;
//     }

//     public static void main(String[] args) {
//         int result = 1;
//         for (int i = 1; i <= MAX_N; i++) {
//             int currentGcd = gcd(i, result);
//             result = result / currentGcd * i;
//         }
//         System.out.println(result);
//     }
// }