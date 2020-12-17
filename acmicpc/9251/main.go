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

    sc.Scan()
    s1 := sc.Bytes()
    sc.Scan()
    s2 := sc.Bytes()
    wt.WriteString(strconv.Itoa(LCS(&s1, &s2)))
}

func LCS(s1, s2 *[]byte) int {
    dp := make([][]int, len(*s2)+1, len(*s2)+1)
    for k := range dp {
        dp[k] = make([]int, len(*s1)+1, len(*s1)+1)
    }
    for k1, v1 := range *s2 {
        for k2, v2 := range *s1 {
            if v1 == v2 {
                dp[k1+1][k2+1] = dp[k1][k2] + 1
            } else {
                dp[k1+1][k2+1] = Max(dp[k1][k2+1], dp[k1+1][k2])
            }
        }
    }
    return dp[len(*s2)][len(*s1)]
}

func Max(a, b int) (max int) {
    max = a
    if a < b {
        max = b
    }
    return
}
