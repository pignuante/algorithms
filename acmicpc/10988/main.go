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
    text := sc.Text()
    r := '1'
    for i, l := 0, len(text); i < l/2; i++ {
        if text[i] != text[l-i-1] {
            r = '0'
            break
        }
    }
    wt.WriteRune(r)
}
