package main

import "testing"

// TestGeneration this function tests several test cases
func TestGeneration(t *testing.T) {
	cases := []struct{ in, expected int }{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
	}
	var fib Fibonacci
	f := fib.Generate()
	for _, c := range cases {
		if result := f(); result != c.expected {
			t.Errorf("Expected %d, got %d", result, c.expected)
		}
	}
}
