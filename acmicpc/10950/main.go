package main

import "fmt"

func main() {
    var test, A, B int

   fmt.Scanf("%d", &test)
   for i := 0; i < test; i++ {
        fmt.Scanf("%d %d", &A, &B)
        fmt.Printf("%d\n", A+B)
   }
}
