package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    nums := make([]int, N, N)
    for k := range nums {
        sc.Scan()
        t, _ := strconv.Atoi(sc.Text())
        nums[k] = t
    }
    sort.Ints(nums)
    for _, k := range nums {
        wt.WriteString(strconv.Itoa(k) + "\n")
    }
}
