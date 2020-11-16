package main

import "fmt"

func main() {
    var grade int
    fmt.Scanf("%d", &grade)

    if grade > 89 {
        fmt.Println("A")
    } else if grade > 79 {
        fmt.Println("B")
    } else if grade > 69 {
        fmt.Println("C")
    } else if grade > 59 {
        fmt.Println("D")
    } else {
        fmt.Println("F")
    }

}
