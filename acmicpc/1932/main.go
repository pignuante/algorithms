package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r = 0

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
    nums := make([][]int, n, n)
    nums[0] = []int{nextInt()}
    for y := 1; y < n; y++ {
        X := make([]int, y+1, y+1)
        for x := 0; x < y+1; x++ {
            t := nextInt()
            if x == 0 {
                X[x] = nums[y-1][x] + t
            } else if y == x {
                X[x] = nums[y-1][x-1] + t
            } else {
                X[x] = Max(nums[y-1][x-1], nums[y-1][x]) + t
            }
        }
        nums[y] = X
    }
    wt.WriteString(strconv.Itoa(Max(nums[n-1]...)))
}

func Max(nums ...int) (max int) {
    max = -1
    for _, v := range nums {
        if max < v {
            max = v
        }
    }
    return
}
