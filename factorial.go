package interview

// Factorial ...
func Factorial(n int64) int64 {

	out := int64(1)

	if n < 0 {
		return 0
	}

	for i := int64(2); i <= n; i++ {
		out *= i
	}

	return out
}

// FactorialRecursive ...
func FactorialRecursive(n int64) int64 {

	if n < 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}
