/*
	A palindromic number reads the same both ways. The largest palindrome made
	from the product of two 2-digit numbers is 9009 = 91 × 99.

	Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "fmt"
import "strconv"

func isPalindrome(n int) bool {
	numstr := strconv.Itoa(n)
	for i, j := 0, len(numstr)-1; i < j; i, j = i+1, j-1 {
		if numstr[i] != numstr[j] {
			return false
		}
	}

	return true
}

func maxPal() int {
	n, max := 0, 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			n = i*j
			if isPalindrome(n) {
				if n > max {
					max = n
				}
			}
		}
	}

	return max
}

func main() {
	fmt.Println("Largest palindrome from the product of two 3-digit numbers: ", maxPal())
}
