package main

import "fmt"

func main() {
    var x, y int
    fmt.Scanf("%d\n%d", &x, &y)

    if x>0 && y>0 {
        fmt.Println(1)
    }
    if x<0 && y>0 {
        fmt.Println(2)
    }
    if x<0 && y<0 {
        fmt.Println(3)
    }
    if x>0 && y<0 {
        fmt.Println(4)
    }
}
