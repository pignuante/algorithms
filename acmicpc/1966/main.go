package main

import (
    "bufio"
    "os"
    "sort"
    // "strconv"
    "fmt"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c-'0')
    }
    return
}

type node struct {
    idx, priority int
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    T := nextInt()
    for ; T > 0; T-- {
        N, M := nextInt(), nextInt()
        que := make([]node, N)
        for i := 0; i < N; i++ {
            que[i] = node{idx:i, priority:nextInt()}
        }
        sort.Slice(que, func(i, j int) (bool) {
            return que[i].priority >  que[j].priority
        })
        fmt.Println(que, M)
    }
}
