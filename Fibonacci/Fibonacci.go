package main

// Fibonacci is an exported type that generates Fibonacci numbers
type Fibonacci struct{}

// Get is a function which return n-th number.
func (f Fibonacci) Get(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return f.Get(n-1) + f.Get(n-2)
}

// Generate is a function which generates consecutive numbers.
func (f Fibonacci) Generate() func() int {
	var a, b int = 0, 1
	return func() int {
		var tmp = a
		a = b
		b = b + tmp
		return tmp
	}
}
