package main

import (
    "bufio"
    "os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    names := sc.Text()
    for _, v := range names {
        if 'A' <= v && v <= 'Z' {
            wt.WriteByte(byte(v))
        }
    }
}
