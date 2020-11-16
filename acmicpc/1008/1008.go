package main

import "fmt"

func main(){
    var A, B float64
    fmt.Scanf("%f %f", &A, &B)
    fmt.Printf("%0.15f", A/B)
}
