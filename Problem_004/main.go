package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	var nDigitNum int
	fmt.Print("Enter how many digit-number should be: ")
	for {
		fmt.Print("")
		_, err := fmt.Scan(&nDigitNum)
		if err != nil {
			log.Fatal(err)
		}
		if nDigitNum > 0 {
			break
		}
		fmt.Println("Please enter number greater then 0.")
	}

	palindrome, factor1, factor2 := largestPalindrome(nDigitNum)
	fmt.Printf("Max palindrome for %d-digit number is: %d (factor1: %d, factor2: %d)", nDigitNum, palindrome, factor1, factor2)
}

func largestPalindrome(nDigit int) (maxPalindrome, maxFactor1, maxFactor2 int) {
	end := int(math.Pow10(nDigit)) - 1
	start := int(math.Pow10(nDigit - 1))

	for i := end; i >= start; i-- {
		for j := end; j >= start; j-- {
			product := i * j
			if maxPalindrome > product {
				break
			}
			if isPalindromeNumber(int64(product)) {
				maxPalindrome = product
				maxFactor1 = i
				maxFactor2 = j
			}
		}
	}
	return maxPalindrome, maxFactor1, maxFactor2
}

func isPalindromeNumber(num int64) bool {
	str := strconv.Itoa(int(num))
	length := len(str)

	for i := 0; length-(2*i) >= 1; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}
