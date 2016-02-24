package main

import "fmt"

func main() {
	var x int
	fmt.Printf("Enter number: ")
	if _, err := fmt.Scanln(&x); err != nil || x < 0 {
		fmt.Println("Invalid number!")
		return
	}

	var fib Fibonacci
	fmt.Printf("n-th number is %d\n\n", fib.Get(x))

	fmt.Println("Fibonacci Numbers:")
	f := fib.Generate()
	for i := 0; i <= x; i++ {
		fmt.Printf("F(%d)=%d\n", i, f())
	}
}
