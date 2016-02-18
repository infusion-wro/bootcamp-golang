package main

import "fmt"

func fibonacci() func() int {
    var a,b int = 0,1;
    return func() int {
        var tmp = a;
        a = b;
        b = b+tmp;
        return tmp;
    };
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