package main

import "fmt"

func Fibonacci(a int) int {
}

func main() {
    var x int
    fmt.Printf("Enter number: ")
    if _,err:=fmt.Scanln(&x);err != nil || x < 0 {
        fmt.Println("Invalid number!")
        return;
    }
    
    f := Fibonacci(x);
    fmt.Println(f);
}