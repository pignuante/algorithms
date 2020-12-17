package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    rd := bufio.NewReader(os.Stdin)
    defer wt.Flush()

    var N int
    fmt.Fscanf(rd, "%d", &N)
    nums := make([]float64, N+1, N+1)

    for i := 0; i <= N; i++ {
        fmt.Fscanf(rd, "%f\n", &nums[i])
    }
    nums = nums[1:]
    wt.WriteString(DP(&nums))
}

func DP(nums *[]float64) string {
    l := len(*nums)
    dp := make([]float64, l, l)
    dp[0] = (*nums)[0]

    max := float64(-1)
    for i := 1; i < l; i++ {
        dp[i] = Max((*nums)[i], (*nums)[i]*dp[i-1])
        if max < dp[i] {
            max = dp[i]
        }
    }
    return fmt.Sprintf("%0.3f", max)
}

func Max(a, b float64) (max float64) {
    max = a
    if a < b {
        max = b
    }
    return
}
