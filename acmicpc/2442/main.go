package main

import (
    "bufio"
    "os"
    // "strconv"
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
        for s := N - i - 1; s > 0; s-- {
            wt.WriteString(" ")
        }
        for x := 0; x < 2*i+1; x++ {
            wt.WriteString("*")
        }
        wt.WriteByte('\n')
    }
}
