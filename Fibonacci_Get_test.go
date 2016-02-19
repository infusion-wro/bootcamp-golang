package main

import "testing"

// TestFirstNumber this function tests whether first number is zero
func TestGetFirstNumber(t *testing.T) {
	var fib Fibonacci
	result := fib.Get(0)
	if result != 0 {
		t.Error("Expected 0, got ", result)
	}
}

// TestSecondNumber this function tests whether second number is one
func TestGetSecondNumber(t *testing.T) {
	var fib Fibonacci
	result := fib.Get(1)
	if result != 1 {
		t.Error("Expected 1, got ", result)
	}
}

// TestOtherCases this function tests several test cases
func TestGetOtherCases(t *testing.T) {
	cases := []struct{ in, expected int }{
		{2, 1},
		{5, 5},
		{10, 55},
		{19, 4181},
	}
	var fib Fibonacci
	for _, c := range cases {
		got := fib.Get(c.in)
		if got != c.expected {
			t.Errorf("Expected %d, got %d", got, c.expected)
		}
	}
}
