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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    dp := []bool{false, false, true, false}
    for i := 4; i <= N; i++ {
        if !dp[i-1] || !dp[i-3] {
            dp = append(dp, true)
        } else {
            dp = append(dp, false)
        }
    }
    if dp[N] {
        wt.WriteString("SK")
    } else {
        wt.WriteString("CY")
    }
}
