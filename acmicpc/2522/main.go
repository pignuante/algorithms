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
        for space := N - i - 1; space > 0; space-- {
            wt.WriteByte(' ')
        }
        for star := 0; star < i+1; star++ {
            wt.WriteByte('*')
        }
        wt.WriteByte('\n')
    }
    for i := 1; i < N; i++ {
        for space := 0; space < i; space++ {
            wt.WriteByte(' ')
        }
        for star := N; star > i; star-- {
            wt.WriteByte('*')
        }
        wt.WriteByte('\n')
    }
}
