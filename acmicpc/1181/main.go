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
    words := make([]string, N, N)
    for ; N > 0; N-- {
        sc.Scan()
        words[N-1] = sc.Text()
    }
    sort.Slice(words, func(i, j int) bool {
        if len(words[i]) == len(words[j]) {
            return words[i] < words[j]
        }
        return len(words[i]) < len(words[j])
    })
    before := ""
    for _, v := range words {
        if before != v {
            wt.WriteString(v + "\n")
            before = v
        }
    }
}
