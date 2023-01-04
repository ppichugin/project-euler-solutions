package main

import (
	"fmt"
	"math"
)

var (
	maxPrime     = 2_000_000
	isPrime      = make(map[int]int, int(math.Sqrt(float64(maxPrime))))
	counter, sum int
)

func main() {
	SieveOfEratosthenes(maxPrime)
	fmt.Printf("The sum of the prime numbers below %d is: %d\n", maxPrime, sum)
}

func SieveOfEratosthenes(value int) {
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			addPrimeToMap(i)
			sum += i
		}
	}
}

func addPrimeToMap(i int) {
	counter++
	isPrime[counter] = i
}
