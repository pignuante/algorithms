package main

import (
    "bufio"
    "os"
    "strconv"
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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    K := nextInt()
    A, B := Click(K)
    wt.WriteString(strconv.Itoa(A))
    wt.WriteByte(' ')
    wt.WriteString(strconv.Itoa(B))
}

func Click(K int) (A, B int) {
    A, B = 0, 1
    for i := 2; i <= K; i++ {
        A, B = B, A+B
    }
    return
}
