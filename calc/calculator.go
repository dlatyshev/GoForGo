package calc

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func Modulus(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a - (b * float64(int(a/b)))
}

func Power(a, b float64) float64 {
	result := 1.0
	for i := 0; i < int(b); i++ {
		result *= a
	}
	return result
}

func SquareRoot(a float64) float64 {
	if a < 0 {
		return 0
	}
	result := a
	for i := 0; i < 10; i++ {
		result = (result + a/result) / 2
	}
	return result
}

func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func Fibonacci(n int) []int {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}
