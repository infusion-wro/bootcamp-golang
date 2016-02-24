package main

import "fmt"

func main() {
	n := 101
	a := ReadArray("missing.json")
	var sum int
	half := len(a) / 2
	/*	TASK 1
		c := make(chan int)
		go sumToChannel(a[:half], c)
		go sumToChannel(a[half:], c)
		x, y := <-c, <-c
		sum = x + y
	*/

	/*	TASK 2
		c1 := gosum(a[:half])
		c2 := gosum(a[half:])
		x, y := <-c1, <-c2
		sum := x + y
	*/
	expected := (n * (1 + n)) / 2
	fmt.Printf("We are missing: %d\n", expected-sum)
}

func sumToChannel(x []int, c chan int) {
	var sum int
	for _, i := range x {
		sum += i
	}
	c <- sum
}

func gosum(x []int) <-chan int {
	c := make(chan int)
	go func() {
		var sum int
		for _, i := range x {
			sum += i
		}
		c <- sum
	}()
	return c
}
