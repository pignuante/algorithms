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

    n := nextInt()

    wt.WriteString(strconv.Itoa(Recurrence(n)))
}

func Recurrence(n int) int {
    dp := make([]int, n+2, n+2)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        for j := 0; j < i; j++ {
            dp[i] += dp[j] * dp[i-j-1]
        }
    }
    return dp[n]
}
