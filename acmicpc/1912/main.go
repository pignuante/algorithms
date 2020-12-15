package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    flag := 1
    for _, c := range sc.Bytes() {
        if c == '-' {
            flag = -1
            continue
        }
        r *= 10
        r += int(c - '0')
    }
    r *= flag
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    n := nextInt()
    nums := make([]int, n, n)
    dp := make([]int, n, n)
    for i := 0; i < n; i++ {
        nums[i] = nextInt()
    }
    dp[0] = nums[0]
    result := dp[0]

    for i := 1; i < n; i++ {
        dp[i] = Max(dp[i-1]+nums[i], nums[i])
        result = Max(result, dp[i])
    }
    wt.WriteString(strconv.Itoa(result))
}

func Max(a, b int) (r int) {
    r = a
    if b > a {
        r = b
    }
    return
}
