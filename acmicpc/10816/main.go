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
    cards := make(map[int]int, N)
    for ; N > 0; N-- {
        sc.Scan()
        t, _ := strconv.Atoi(sc.Text())
        cards[t]++

    }
    sc.Scan()
    M, _ := strconv.Atoi(sc.Text())
    for ; M > 0; M-- {
        sc.Scan()
        t, _ := strconv.Atoi(sc.Text())
        if v, exists := cards[t]; exists {
            wt.WriteString(strconv.Itoa(v) + " ")
        } else {
            wt.WriteString("0 ")
        }
    }
}
