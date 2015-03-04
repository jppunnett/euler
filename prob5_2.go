/*
	2520 is the smallest number that can be divided by each of the numbers 
	from 1 to 10 without any remainder.
	
	What is the smallest positive number that is evenly divisible by all of the 
	numbers from 1 to 20?
*/
package main

import (
	"fmt"
	"time"
	"math"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19}

type PrimeFactor struct {
	prime, exponent int
}

// Find the prime factors of n. Only looks at primes between 2-20.
func primeFactors(n int) []PrimeFactor {

	pfs := []PrimeFactor{}
	var exp int

	for _, p := range primes {
		exp = 0
		for ; n != 0 && n%p == 0; n/=p {
			exp++
		}

		if exp != 0 {
			pfs = append(pfs, PrimeFactor{p, exp})
		}

	}

	return pfs
}

func main() {

	t0 := time.Now()

	mpfs := map[int]PrimeFactor{}

	for i := 2; i <= 20; i++ {
		pfs := primeFactors(i)
		for _, pf := range pfs {
			if elem, ok := mpfs[pf.prime]; ok {
				if pf.exponent > elem.exponent  {
					mpfs[pf.prime] = pf
				}
			} else {
				mpfs[pf.prime] = pf
			}
		}
	}

	n := 1
	for _, pf := range mpfs {
		n *= int(math.Pow(float64(pf.prime), float64(pf.exponent)))
	}

	fmt.Println(n)

	t1 := time.Since(t0)
	fmt.Printf("Took %v to run.\n", t1)
}
