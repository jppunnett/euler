/*
	Find the difference between the sum of the squares of the first one hundred
	natural numbers and the square of the sum.
*/
package main

import (
	"fmt"
	"time"
)

func bruteForce(n uint) uint {
	var i uint

	// Sum of squares 1..n
	var sumsqr uint = 0
	for i = 1; i <= n; i++ {
		sumsqr += i*i
	}
	
	// Square of sum 1..n
	var sqrsum uint = 0
	for i = 1; i <= n; i++ {
		sqrsum += i
	}
	sqrsum *= sqrsum

	return sqrsum-sumsqr
}

func formula(n float32) uint {
	// Sum of squares 1..n
	sumsqr := (n*n*n/3) + (n*n/2) + (n/6)

	// Square of sum 1..n
	sqrsum := n*(n+1)/2
	sqrsum *= sqrsum

	return uint(sqrsum-sumsqr)
}

func main() {

	t0 := time.Now()
	n1 := bruteForce(100)
	t1 := time.Since(t0)

	fmt.Println(n1)
	fmt.Printf("Brute force approach took %v to run.\n", t1)

	t0 = time.Now()
	n2 := formula(100)
	t1 = time.Since(t0)

	fmt.Println(n2)
	fmt.Printf("Formula approach took %v to run.\n", t1)
}
