package main

import . "github.com/hortencio-main/uint128"

import "fmt"

func main() {
    var a, b Uint128 
    
    a = a.Add64(1)
    b = b.Add64(3)
    
    // calculate the 50th power of 3
    for i := 0; i < 50; i++ {
        a = a.Mul128(b)
        fmt.Printf("%s\n", a.String())
    }
    
}
