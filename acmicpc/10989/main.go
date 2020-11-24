package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    nums := make([]int, 10000+1)
    for ; N > 0; N-- {
        sc.Scan()
        n, _ := strconv.Atoi(sc.Text())
        nums[n]++
    }
    for k, v := range nums {
        if v != 0 {
            for i := 0; i < v; i++ {
                wt.WriteString(strconv.Itoa(k) + "\n")
            }
        }
    }
}
