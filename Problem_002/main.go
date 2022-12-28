package main

import "fmt"

// Solution is based on the fact that Fibonacci numbers has pattern:
// odd, odd, even, odd, odd, even, odd, odd, even…
// 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
// sum of first and second even Fibonacci numbers:
// 2 + 8 = 10, i.e. :
// 4 x (Fn-1) + (Fn-2)
func main() {
	fmt.Println("\n-----------------------------")
	n := 35
	fmt.Printf("Sum of all even Fibonacci numbers that do not exceed (%d): %d\n", n, sumEvenFibonacci(n)) // 44
	fmt.Println("\n-----------------------------")
	n = 4_000_000
	fmt.Printf("Sum of all even Fibonacci numbers that do not exceed (%d): %d\n", n, sumEvenFibonacci(n)) // 4613732
}

func sumEvenFibonacci(n int) (sum int) {
	sum = 10
	gen := makeEvenFibGen()

	for sum < n {
		sum += gen()
	}
	return sum
}

func makeEvenFibGen() func() int {
	fn_2 := 2 // Fn-2
	fn_1 := 8 // Fn-1
	QtyNums := 2
	fmt.Printf("Even n-th Fib num: Fn(1-st): \t%d\n", fn_2)
	fmt.Printf("Even n-th Fib num: Fn(2-nd): \t%d\n", fn_1)
	return func() int {
		var fn int
		fn = 4*fn_1 + fn_2
		fn_2, fn_1 = fn_1, fn
		QtyNums++
		fmt.Printf("Even n-th Fib num: Fn(%d-th): \t%d\n", QtyNums, fn)
		return fn_1
	}
}
