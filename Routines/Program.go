package main

import "fmt"

func main() {
	n := 101
	a := ReadArray("missing.json")
	var sum int

	// put content here :-)

	expected := (n * (1 + n)) / 2
	fmt.Println(a)
	fmt.Printf("We are missing: %d\n", expected-sum)
}

// sumToChannel is a function that will sum input and save result into channel
// hint: this is synchrounous function, will be called with 'go' keyword
func sumToChannel(x []int, c chan int) {
}

// gosum is function that returns channel with sum (of input array)
// hint: this function will trigger async computations and return channel
func gosum(x []int) <-chan int {
	return make(chan int)
}
