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
    A := make([]int, N, N)
    dp := make([]int, N, N)
    for i := 0; i < N; i++ {
        A[i] = nextInt()
    }
    dp[0] = 1
    max := 1
    for i := 1; i < N; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if A[i] > A[j] && dp[j]+1 > dp[i] {
                dp[i] = dp[j] + 1
                if dp[i] > max {
                    max = dp[i]
                }
            }
        }
    }
    wt.WriteString(strconv.Itoa(max))
}
