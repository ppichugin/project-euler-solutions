package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Solution O(1) ---")
	fmt.Println("Multiples below 10: \t", sumOfMultiplies(10))
	fmt.Println("Multiples below 1000: \t", sumOfMultiplies(1000))
	fmt.Println("--- Solution O(n) ---")
	fmt.Println("Multiples below 10: \t", sumOfMultiplesInCycle(10))
	fmt.Println("Multiples below 1000: \t", sumOfMultiplesInCycle(1000))
}

func sumOfMultiplies(n int) int {
	var (
		LIMIT           = n - 1
		upperForThree   = LIMIT / 3
		upperForFive    = LIMIT / 5
		upperForFifteen = LIMIT / 15
	)

	// calculate sums
	sumThree := 3 * upperForThree * (1 + upperForThree) / 2
	sumFive := 5 * upperForFive * (1 + upperForFive) / 2
	sumFifteen := 15 * upperForFifteen * (1 + upperForFifteen) / 2

	return sumThree + sumFive - sumFifteen
}

func sumOfMultiplesInCycle(n int) int {
	var sum int
	for i := 1; i < n; i++ {
		num := i
		if num%3 == 0 || num%5 == 0 {
			sum += num
		}
	}
	return sum
}
