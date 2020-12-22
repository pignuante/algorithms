package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

func main() {
    const N = 99
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    n, k := nextInt(), nextInt()
    cards := make([]int, n, n)
    results := make(map[string]bool)
    visited := make([]bool, N+1, N+1)
    for i := 0; i < n; i++ {
        cards[i] = nextInt()
    }
    DFS(&cards, 0, k, "", &results, &visited)
    wt.WriteString(strconv.Itoa(len(results)))
}

func DFS(cards *[]int, idx, k int, str string, results *map[string]bool, visited *[]bool) {
    if idx == k {
        (*results)[str] = true
        return
    }
    for i, card := range *cards {
        if (*visited)[i] {
            continue
        }
        (*visited)[i] = true
        DFS(cards, idx+1, k, str+strconv.Itoa(card), results, visited)
        (*visited)[i] = false
    }
}
