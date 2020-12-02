package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanWords)

    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    A := make(map[string]int, N)
    for i := 0; i < N; i++ {
        sc.Scan()
        t := sc.Text()
        A[t] = 1
    }
    sc.Scan()
    M, _ := strconv.Atoi(sc.Text())
    for ; M > 0; M-- {
        sc.Scan()
        t := sc.Text()
        if A[t] == 1 {
            wt.WriteString("1\n")
        } else {
            wt.WriteString("0\n")
        }
    }
}
