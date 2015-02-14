// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?
package main

import (
	"fmt"
	"math"
)

func maxPrimeFactor(n uint64) uint64 {
	
	var primeFactors []uint64

	for ; n%2 == 0; n/=2 {
		primeFactors = append(primeFactors, 2)
	}

	fltn := float64(n)
	sqrtn := uint64(math.Sqrt(fltn))

	var i uint64 = 3
	for ; i < sqrtn; i+=2 {
		for ; n%i == 0; n/=i {
			primeFactors = append(primeFactors, i)
		}
	}

	if n > 2 {
		primeFactors = append(primeFactors, n)
	}

	return primeFactors[len(primeFactors)-1]
}

func main() {
	fmt.Println(maxPrimeFactor(600851475143))
}
