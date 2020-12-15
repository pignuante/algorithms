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
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, M := nextInt(), nextInt()
    visited := make([]bool, N+1, N+1)
    points := make([][]int, N+1, N+1)
    for i := 0; i < M; i++ {
        u, v := nextInt(), nextInt()
        points[u] = append(points[u], v)
        points[v] = append(points[v], u)
    }

    count := 0
    for i := 1; i <= N; i++ {
        if !visited[i] {
            DFS(&points, &visited, i)
            count++
        }
    }
    wt.WriteString(strconv.Itoa(count))
}

func DFS(points *[][]int, visited *[]bool, idx int) {
    (*visited)[idx] = true
    for _, next := range (*points)[idx] {
        if !(*visited)[next] {
            DFS(points, visited, next)
        }
    }
}
