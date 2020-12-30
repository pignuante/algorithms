package main

import (
    "bufio"
    "os"
    "sort"
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

    N, max := nextInt(), -1
    lopes := make([]int, N, N)
    for i := 0; i < N; i++ {
        lopes[i] = nextInt()
    }
    sort.Ints(lopes)
    for i, lope := range lopes {
        max = Max(max, lope*(N-i))
    }
    wt.WriteString(strconv.Itoa(max))
}

func Max(a, b int) (max int) {
    max = a
    if a < b {
        max = b
    }
    return
}
