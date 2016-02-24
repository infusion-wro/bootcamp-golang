package main

import "testing"

// TestGeneration this function tests several test cases
func TestGeneration(t *testing.T) {
	numbers := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	var fib Fibonacci
	f := fib.Generate()
	for _, number := range numbers {
		if result := f(); result != number {
			t.Errorf("Expected %d, got %d", number, result)
		}
	}
}
