package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var str string
    reader := bufio.NewReader(os.Stdin)
    fmt.Fscanln(reader, &str)

    var times = []int{
        3, 3, 3,
        4, 4, 4,
        5, 5, 5,
        6, 6, 6,
        7, 7, 7,
        8, 8, 8, 8,
        9, 9, 9,
        10, 10, 10, 10}
        var result int
        for _, v := range str {
            result += times[int(v)-65]
        }
        fmt.Println(result)
    }
