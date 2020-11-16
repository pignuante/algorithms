package main

import "fmt"

func main() {
    var A,B,C int
    fmt.Scanf("%d %d %d", &A,&B,&C)

    fmt.Println((A+B)%C)
    fmt.Println(((A%C) + (B%C))%C)
    fmt.Println((A*B)%C)
    fmt.Println(((A%C) * (B%C))%C)
}
