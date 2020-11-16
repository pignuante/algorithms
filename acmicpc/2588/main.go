package main

import "fmt"

func main() {
    var a, b int
    
    fmt.Scanf("%d\n%d", &a, &b)
    
    temp := b
    for temp != 0 {
        t := temp % 10
        fmt.Println(a*t)
        temp /= 10
    }
    fmt.Println(a*b)

}
