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
    dp := make([][]int, n, n)
    dp[0] = []int{1}

    Pascal(&dp)
    wt.WriteString(strconv.Itoa(dp[n-1][k-1]))
}

func Pascal(dp *[][]int) {
    n := len(*dp)
    for y := 1; y < n; y++ {
        for x := 0; x < y+1; x++ {
            if x == 0 || x == y {
                (*dp)[y] = append((*dp)[y], 1)
            } else {
                (*dp)[y] = append((*dp)[y], (*dp)[y-1][x-1]+(*dp)[y-1][x])
            }
        }
    }
}
