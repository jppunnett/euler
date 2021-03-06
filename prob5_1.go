/*
	2520 is the smallest number that can be divided by each of the numbers 
	from 1 to 10 without any remainder.
	
	What is the smallest positive number that is evenly divisible by all of the 
	numbers from 1 to 20?
*/
package main

import "fmt"
import "time"

const k int = 20

func isEvenlyDivisible(n int) bool {
	for i := 2; i <= k; i++ {
		if n%i != 0 {
			return false
		}
	}

	return true	
}

func main() {

	t0 := time.Now()

	i := k
	for ; true; i+=k {
		if isEvenlyDivisible(i) {
			break;
		}
	}

	t1 := time.Now()
	
	fmt.Println(i)
	fmt.Printf("Took %v to run.\n", t1.Sub(t0))
}
