package main

import "fmt"

func Fibonacci(a int) int {
    if(a == 0){
        return 0;
    }
    if(a == 1){
        return 1;
    }
    
    return Fibonacci(a-1) + Fibonacci(a-2);
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