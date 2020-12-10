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

    for y := 0; y < N; y++ {
        for space := 0; space < N-y-1; space++ {
            wt.WriteByte(' ')
        }
        for star := 0; star < y+1; star++ {
            wt.WriteByte('*')
            wt.WriteByte(' ')
        }
        wt.WriteByte('\n')
    }
}
