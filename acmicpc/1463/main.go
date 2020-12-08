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
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    dp := make([]int, N+1, N+1)
    for k := range dp[:N] {
        dp[k] = 1000000
    }
    dp[N-1] = 1
    for i := N; i > 1; i-- {
        if i%3 == 0 {
            dp[i/3] = Min(dp[i]+1, dp[i/3])
        }
        if i%2 == 0 {
            dp[i/2] = Min(dp[i]+1, dp[i/2])
        }
        dp[i-1] = Min(dp[i]+1, dp[i-1])
    }
    wt.WriteString(strconv.Itoa(dp[1]))
}

func Min(a, b int) (min int) {
    min = b
    if a < b {
        min = a
    }
    return
}
