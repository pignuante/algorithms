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

    stairs := nextInt()
    points := make([]int, stairs+3, stairs+3)
    dp := make([]int, stairs+3, stairs+3)
    for i := 3; i < stairs+3; i++ {
        points[i] = nextInt()
        dp[i] = Max(dp[i-2], dp[i-3]+points[i-1]) + points[i]
    }
    wt.WriteString(strconv.Itoa(dp[stairs+2]))
}

func Max(nums ...int) (max int) {
    max = 0
    for _, v := range nums {
        if max < v {
            max = v
        }
    }
    return
}
