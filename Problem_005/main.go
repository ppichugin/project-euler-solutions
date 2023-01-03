package main

// Euclidean algorithm
// a x b = LCM(a, b) * GCD(a, b)
// LCM(a, b) = (a x b) / GCD(a, b)

func main() {
	lcmNum := 1
	end := 20
	for i := 2; i <= end+1; i++ {
		lcmNum = lcm(i, lcmNum)
	}
	println(lcmNum)
}

// Greatest common divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Least common multiple
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
