package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    chars := 0
    for y := 0; y < 8; y++ {
        sc.Scan()
        flag := y % 2
        for k, v := range sc.Bytes() {
            if k%2 == flag && v == 'F' {
                chars++
            }
        }
    }
    wt.WriteString(strconv.Itoa(chars))
}
