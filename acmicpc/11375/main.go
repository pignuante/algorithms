package main

import (
    "bufio"
    "os"
    "strconv"
    "fmt"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c-'0')
    }
    return
}

var (
    N, M int
    Works [][]int
    Visited []bool
    Match  []int
)
func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, M = nextInt(), nextInt()
    GetSchedule()
    Visited, Match = make([]bool, N+1, N+1), make([]int, M+1, M+1)
    fmt.Println(len(Visited), Visited)
    count := 0
    for i := 1; i <= N; i++ {
        for j := range Visited {
            Visited[j] = false
        }
        fmt.Println(Visited)
        fmt.Println(Match)
        if Bipartite(i) {
            count++
        }
    }
    wt.WriteString(strconv.Itoa(count))
}

func Bipartite(person int) (ok bool) {
    fmt.Println("?",person)
    ok = false
    if !Visited[person] {
        Visited[person] = true
        for _, work := range Works[person] {
            if Match[work] == 0 || Bipartite(Match[work]) {
                fmt.Println(person, work)
                Match[work] = person
                ok = true
            }
        }
    }
    return
}

func GetSchedule() {
    Works = make([][]int, N+1, N+1)
    for i := 1; i <= N; i++ {
        w := nextInt()
        Works[i] = make([]int, w, w)
        for j := 0; j < w; j++ {
            Works[i][j] = nextInt()
        }
    }
}
