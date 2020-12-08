package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r, _ = strconv.Atoi(sc.Text())
    return
}
func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    prices := make([][]int, N, N)
    prices[0] = []int{nextInt(), nextInt(), nextInt()}

    for i := 1; i < N; i++ {
        prices[i] = []int{nextInt(), nextInt(), nextInt()}
        prices[i][0] = Min(prices[i-1][1], prices[i-1][2]) + prices[i][0]
        prices[i][1] = Min(prices[i-1][0], prices[i-1][2]) + prices[i][1]
        prices[i][2] = Min(prices[i-1][0], prices[i-1][1]) + prices[i][2]

    }
    wt.WriteString(strconv.Itoa(Min(prices[N-1]...))+"\n")
}

func Min(nums ...int) (min int) {
    min = 1000001
    for _, v := range nums {
        if min > v {
            min = v
        }
    }
    return
}
