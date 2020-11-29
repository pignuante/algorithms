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
    NK := strings.Fields(sc.Text())
    N, _ := strconv.Atoi(NK[0])
    K, _ := strconv.Atoi(NK[1])
    j := make([]int, N)
    for k := range j {
        j[k] = k + 1
    }
    idx := K - 1
    wt.WriteString("<")
    for len(j) > 1 {
        wt.WriteString(strconv.Itoa(j[idx]) + ", ")
        j = append(j[:idx], j[idx+1:]...)
        idx = (idx + K - 1) % len(j)
    }
    wt.WriteString(strconv.Itoa(j[0]) + ">")
}
