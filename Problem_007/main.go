package main

import (
	"fmt"
	"math"
	"math/big"
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
			addPrimeToMap(i)
		}
	}
}

func addPrimeToMap(i int) {
	counter++
	isPrime[counter] = i
}

func main() {
	var nthNum int
	_, err := fmt.Scan(&nthNum)
	if err != nil {
		return
	}
	fmt.Println("Solution with Sieve Of Eratosthenes")
	SieveOfEratosthenes(maxEdge)
	fmt.Printf("The %dth prime number is: %d\n", nthNum, isPrime[nthNum])

	// ProbablyPrime is 100% accurate for inputs less than 2⁶⁴.
	// 18_446_744_073_709_551_616 - which is quite enough for our task.
	fmt.Println("Solution with the Miller-Rabin test")
	counter = 0
	isPrime = make(map[int]int, int(math.Sqrt(float64(maxEdge))))
	for i := 2; ; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			addPrimeToMap(i)
		}
		if counter == nthNum {
			break
		}
	}
	fmt.Printf("The %dth prime number is: %d\n", nthNum, isPrime[nthNum])
}
