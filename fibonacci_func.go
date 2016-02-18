package main

import "fmt"

func fibonacci() func() int {
}

func main() {
    var x int
    fmt.Printf("Enter number: ")
    if _,err:=fmt.Scanln(&x);err != nil {
        fmt.Println("Not a number!")
    }
    
    f := fibonacci()
	for i := 0; i < x; i++ {
		fmt.Println(f())
	}
}