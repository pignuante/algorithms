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
    wt.WriteString(strconv.Itoa(SquareNumber(n)))
}

func SquareNumber(n int) int {
    dp := make([]int, n+1, n+1)
    dp[0], dp[1] = 0, 1
    for i := 1; i <= n; i++ {
        min := 50000*50000 + 1
        for j := 1; j*j <= i; j++ {
            if t := dp[i-j*j]; min > t {
                min = t
            }
        }
        dp[i] = min + 1
    }
    return dp[n]
}
