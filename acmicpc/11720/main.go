package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())

    sc.Scan()
    nums := sc.Text()
    sum := 0
    for i := 0; i < N; i++ {
        num, _ := strconv.Atoi(string(nums[i]))
        sum += num
    }
    wt.WriteString(strconv.Itoa(sum))
}
