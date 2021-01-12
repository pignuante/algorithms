package main

import (
    "bufio"
    "os"
    "strconv"
)
var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
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
    defer wt.Flush()

    T := nextInt()
    for i := 0; i < T; i++ {
        N, M := nextInt(), nextInt()
        for j := 0; j < M; j++ {
            _, _ = nextInt(), nextInt()
        }
        wt.WriteString(strconv.Itoa(N - 1))
        wt.WriteByte('\n')
    }
}
