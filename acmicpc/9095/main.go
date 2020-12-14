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
    dp := make([]int, 11, 11)
    dp[1], dp[2], dp[3] = 1, 2, 4

    T := nextInt()
    for ; T > 0; T-- {
        n := nextInt()
        for y := 4; y <= n; y++ {
            dp[y] = dp[y-1] + dp[y-2] + dp[y-3]
        }
        wt.WriteString(strconv.Itoa(dp[n]) + "\n")
    }
}
