package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Node struct {
    point  int
    edges  []int
    status int
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()
    sc.Scan()
    NMV := strings.Fields(sc.Text())
    N, _ := strconv.Atoi(NMV[0])
    M, _ := strconv.Atoi(NMV[1])
    V, _ := strconv.Atoi(NMV[2])

    nodes := make([]Node, N+1, N+1)
    for i := 1; i <= M; i++ {
        sc.Scan()
        t := strings.Fields(sc.Text())
        p, _ := strconv.Atoi(t[0])
        q, _ := strconv.Atoi(t[1])
        nodes[p].edges = append(nodes[p].edges, q)
        nodes[q].edges = append(nodes[q].edges, p)
    }
    for i := 0; i < N; i++ {
        sort.Ints(nodes[i].edges)
        nodes[i].status = 0
    }
    stack := []int{V}
    DFS(nodes, V, &stack)

    for i := 0; i <= N; i++ {
        sort.Ints(nodes[i].edges)
        nodes[i].status = 0
    }
    visited := []int{}
    BFS(nodes, V, &visited)
    for _, v := range stack {
        wt.WriteString(strconv.Itoa(v)+" ")
    }
    wt.WriteString("\n")
    for _, v := range visited {
        wt.WriteString(strconv.Itoa(v)+" ")
    }
}

func DFS(nodes []Node, start int, stack *[]int) {
    nodes[start].status = 1
    for _, edge := range nodes[start].edges {
        if nodes[edge].status == 0 {
            *stack = append(*stack, edge)
            // node[edge].status = 1
            DFS(nodes, edge, stack)
        }
    }
    nodes[start].status = 2
}

func BFS(nodes []Node, start int, visited *[]int) {
    queue := []int{start}
    nodes[start].status = 1
    for len(queue) > 0 {
        for _, edge := range nodes[queue[0]].edges {
            if nodes[edge].status == 0 {
                nodes[edge].status = 1
                queue = append(queue, edge)
            }
        }
        *visited = append(*visited, queue[0])
        queue = queue[1:]
    }
}
