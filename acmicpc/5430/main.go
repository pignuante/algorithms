package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    tree := make([]int, N+1, N+1)


}
