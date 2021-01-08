package main

import (
    "bufio"
    "os"
    "strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

var (
    N, M    int
    Works   [][]int
    Visited [1001]bool
    Matched [1001]int
)

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, M = nextInt(), nextInt()
    Works = make([][]int, N+1, N+1)
    // Matched = make([]int, M+1, M+1)
    for i := 1; i <= N; i++ {
        w := nextInt()
        Works[i] = make([]int, w+1, w+1)
        for j := 1; j <= w; j++ {
            Works[i][j] = nextInt()
        }
    }
    count := 0
    for i := 1; i <= N; i++ {
        for j := range Visited {
            Visited[j] = false
        }
        if Matching(i) {
            count++
        }
    }
    wt.WriteString(strconv.Itoa(count))
}

func Matching(person int) bool {
    if Visited[person] {
        return false
    }
    Visited[person] = true
    for _, work := range Works[person][1:] {
        if Matched[work] == 0 || Matching(Matched[work]) {
            Matched[work] = person
            return true
        }
    }
    return false
}
