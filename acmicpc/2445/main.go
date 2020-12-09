package main

import (
    "bufio"
    "os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    for i := 0; i < N; i++ {
        for start := 0; start < i+1; start++ {
            wt.WriteByte('*')
        }
        for space := 0; space < N*2-(i+1)*2; space++ {
            wt.WriteByte(' ')
        }
        for start := 0; start < i+1; start++ {
            wt.WriteByte('*')
        }
        wt.WriteByte('\n')
    }
    for i := N - 2; i >= 0; i-- {
        for start := 0; start < i+1; start++ {
            wt.WriteByte('*')
        }
        for space := 0; space < N*2-(i+1)*2; space++ {
            wt.WriteByte(' ')
        }
        for start := 0; start < i+1; start++ {
            wt.WriteByte('*')
        }
        wt.WriteByte('\n')
    }
}
