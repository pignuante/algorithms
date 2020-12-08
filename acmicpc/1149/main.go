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
    for i := 0; i < N; i++ {
        prices[i] = make([]int, 3, 3)
        prices[i] = []int{nextInt(), nextInt(), nextInt()}
    }
    getPrice(&prices)
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
func getPrice(RGBs *[][]int) {
    for i := 1; i < len(*RGBs); i++ {
        (*RGBs)[i][0] = Min((*RGBs)[i-1][1], (*RGBs)[i-1][2]) + (*RGBs)[i][0]
        (*RGBs)[i][1] = Min((*RGBs)[i-1][0], (*RGBs)[i-1][2]) + (*RGBs)[i][1]
        (*RGBs)[i][2] = Min((*RGBs)[i-1][0], (*RGBs)[i-1][1]) + (*RGBs)[i][2]
    }
}
