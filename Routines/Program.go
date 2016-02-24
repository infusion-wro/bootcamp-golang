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

func sumToChannel(x []int, c chan int) {
}

func gosum(x []int) <-chan int {
	return make(chan int)
}
