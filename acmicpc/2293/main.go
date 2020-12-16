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

    n, k := nextInt(), nextInt()
    coins := make([]int, n+1, n+1)
    dp := make([]int, k+1, k+1)
    dp[0] = 1
    for i := 1; i <= n; i++ {
        coins[i] = nextInt()
    }
    DP(&coins, &dp, k)
    wt.WriteString(strconv.Itoa(dp[k]))
}

func DP(coins, dp *[]int, k int) {
    for i := 1; i < len(*coins); i++ {
        for coin := (*coins)[i]; coin <= k; coin++ {
            (*dp)[coin] = (*dp)[coin] + (*dp)[coin-(*coins)[i]]
        }
    }
}
