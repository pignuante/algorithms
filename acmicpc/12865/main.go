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

type Item struct {
    Weight, Value int
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, K := nextInt(), nextInt()
    items := make([]Item, N+1, N+1)
    dp := make([][]int, N+1, N+1)
    for i := 0; i <= N; i++ {
        dp[i] = make([]int, K+1, K+1)
    }
    for i := 1; i <= N; i++ {
        items[i] = Item{Weight: nextInt(), Value: nextInt()}
    }

    for i := 1; i <= N; i++ {
        for j := 1; j <= K; j++ {
            dp[i][j] = dp[i-1][j]
            if w := items[i].Weight; j >= w {
                dp[i][j] = Max(dp[i][j], dp[i-1][j-w]+items[i].Value)
            }
        }
    }
    wt.WriteString(strconv.Itoa(dp[N][K]))
}

func Max(a, b int) (max int) {
    max = a
    if a < b {
        max = b
    }
    return
}
