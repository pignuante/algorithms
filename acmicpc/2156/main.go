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
    wines := make([]int, n+1, n+1)
    for i := 1; i <= n; i++ {
        wines[i] = nextInt()
    }
    dp := drinkWines(&wines)
    wt.WriteString(strconv.Itoa(dp[n]))
}

func drinkWines(wines *[]int) (dp []int) {
    l := len(*wines)
    dp = make([]int, l, l)

    dp[1] = (*wines)[1]
    if l > 2 {
        dp[2] = (*wines)[1] + (*wines)[2]
    }
    for i := 3; i < l; i++ {
        dp[i] = Max(
            (*wines)[i] + (*wines)[i-1] + dp[i-3],
            (*wines)[i] + dp[i-2],
            dp[i-1])
    }
    return
}

func Max(a, b, c int) int {
    if a > b && a > c {
        return a
    } else if b > c {
        return b
    }
    return c
}
