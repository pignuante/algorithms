package main

import (
    "bufio"
    "os"
    "strconv"
    // "fmt"
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

    day := nextInt()
    schedules := make([][2]int, day+2, day+2)
    for i := 1; i <= day; i++ {
        schedules[i] = [2]int{nextInt(), nextInt()}
    }
    wt.WriteString(strconv.Itoa(Working(&schedules, day)))
}

func Working(schedules *[][2]int, day int) int {
    N := len(*schedules)
    dp := make([]int, N, N)
    for i := N - 2; i > 0; i-- {
        if next := i + (*schedules)[i][0]; next > N-1 { // 날짜 over
            dp[i] = dp[i+1]
        } else {
            dp[i] = Max(dp[i+1], dp[i+(*schedules)[i][0]]+(*schedules)[i][1])
        }

    }
    return dp[1]
}

func Max(a, b int) (max int) {
    max = a
    if b > a {
        max = b
    }
    return
}
