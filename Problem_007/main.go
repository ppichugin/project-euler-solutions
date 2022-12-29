package main

import (
	"fmt"
	"math"
)

var (
	// as per Prime Pi function - it says that:
	// There are 10,024 primes less than or equal to 105,000.
	// Which is enough for our calculation
	maxEdge = 105000

	isPrime = make(map[int]int, int(math.Sqrt(float64(maxEdge))))
	counter int
)

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
			counter++
			isPrime[counter] = i
		}
	}
}

func main() {
	var nThNum int
	_, err := fmt.Scan(&nThNum)
	if err != nil {
		return
	}
	SieveOfEratosthenes(maxEdge)
	fmt.Println(isPrime[nThNum])
}
