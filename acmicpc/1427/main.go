package main

import (
    "fmt"
)

func main() {
    var N string
    fmt.Scanf("%s", &N)
    fmt.Println(N)
    for i := '9'; i >= '0'; i-- {
        for _, v := range N {
            if i == v {
                fmt.Printf("%s", string(v))
            }
        }
    }
}
