package main

// Fibonacci is an exported type that generates Fibonacci numbers
type Fibonacci struct{}

// Get is a function which return n-th fibonacci number
func (f Fibonacci) Get(n int) int {
	return 0
}

// Generate is a function which generates consecutive numbers.
func (f Fibonacci) Generate() func() int {
	return func() int {
		return 0
	}
}
