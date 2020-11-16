package main

import "fmt"

func main() {
    var number int
    var result int = 1

    fmt.Scanf("%d", &number)

    for i := 2; i <= number; i++ {
        result += i
    }
    fmt.Printf("%d", result)
}
