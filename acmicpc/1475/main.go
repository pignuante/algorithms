package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    sc.Scan()
    N := sc.Text()
    nums := make([]int, 9, 9)
    for _, v := range N {
        t := v - '0'
        if t == 9 {
            t = 6
        }
        nums[t]++
    }
    nums[6] = (nums[6]+1) / 2
    max := 0
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    wt.WriteString(strconv.Itoa(max))
}
