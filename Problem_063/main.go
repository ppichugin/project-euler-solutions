package main

import (
	"fmt"
	"math"
)

func main() {

	// n cannot be greater than 10 since 10^n is always n+1 digits long.
	fmt.Println("n-digit positive integers:", nDigitsPositiveInteger(9))
}

func nDigitsPositiveInteger(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		in := int(1 / (1 - math.Log10(float64(i))))
		fmt.Printf("Thus %d^n is n digits long when n ≤ %d\n", i, in)
		sum += in
	}
	return sum
}
