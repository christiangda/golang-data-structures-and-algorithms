package interview

// Fibonacci ...
func Fibonacci(n int64) int64 {
	var a int64 = 0
	var b int64 = 1

	for i := int64(0); i < n; i++ {
		b = b + a
		a = b - a
	}

	return a
}

// FibonacciRecursive ...
func FibonacciRecursive(n int64) int64 {

	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)

}
