package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    r := bufio.NewReader(os.Stdin)
    arr := make([]int ,42)
    for i := 0; i< 10; i++ {
        var temp int
        fmt.Fscan(r,&temp)
        arr[temp%42]++
    }
    res := 0
    for _, v := range arr {
        if v != 0 {
            res++
        }
    }
    fmt.Println(res)
}
