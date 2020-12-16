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

    N := nextInt()
    dp := make([][10]int, N+1, N+1)
    for k := range dp[1] {
        dp[1][k] = 1
    }
    for i := 2; i <= N; i++ {
        for left := 0; left < 10; left++ {
            for right := left; right < 10; right++ {
                dp[i][left] = (dp[i][left] + dp[i-1][right]) % 10007
            }
        }
    }
    r := 0
    for _, v := range dp[N] {
        r = (v + r) % 10007
    }
    wt.WriteString(strconv.Itoa(r))
}
