package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

const myNumber = 600851475143

func main() {
	if big.NewInt(myNumber).ProbablyPrime(0) {
		primeFactorFound(myNumber)
	}
	var step = 1
	for i := int(math.Sqrt(myNumber)); i >= 2; i = i - step {
		primeFactor := big.NewInt(int64(i)).ProbablyPrime(0)
		if !primeFactor {
			step = 1
			continue
		} else {
			step = 2
		}
		if myNumber%i == 0 {
			primeFactorFound(i)
		}
	}
}

func primeFactorFound(i int) {
	fmt.Printf("The largest prime factor of the number %d is %d\n", myNumber, i)
	os.Exit(1)
}
