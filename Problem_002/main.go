package main

import "fmt"

// Solution is based on the fact that Fibonacci numbers has pattern:
// odd, odd, even, odd, odd, even, odd, odd, even…
// 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
// sum of first and second even Fibonacci numbers:
// 2 + 8 = 10, i.e. :
// 4 x (Fn-1) + (Fn-2)
func main() {
	fmt.Println(sumEvenFibonacci(35))
	fmt.Println(sumEvenFibonacci(4_000_000))
}

func sumEvenFibonacci(n int) (sum int) {
	sum = 10
	fst := 2 // Fn-2
	snd := 8 // Fn-1
	for {
		fNum := 4*snd + fst
		if fNum >= n {
			return
		}
		sum += fNum
		fst, snd = snd, fNum
	}
}
