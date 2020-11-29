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
    M, _ := strconv.Atoi(NK[1])
    maxSize := M*2 + 1
    if N > M {
        maxSize = N*2 + 1
    }
    location := make([]int, maxSize)
    sec := BFS(location, N, M, maxSize)
    wt.WriteString(strconv.Itoa(sec))
}

func BFS(location []int, from, to, maxSize int) (second int) {
    queue := []int{from}
    location[from] = 1
    for len(queue) > 0 {
        from = queue[0]
        queue = queue[1:]
        if from == to {
            second = location[from]
            break
        }
        Move(location, &queue, from, from*2, maxSize)
        Move(location, &queue, from, from+1, maxSize)
        Move(location, &queue, from, from-1, maxSize)
    }
    return second - 1
}

func Move(location []int, queue *[]int, from, to, maxSize int) {
    if 0 <= to && to < maxSize && (location[to] == 0 || location[from]+1 < location[to]) {
        location[to] = location[from] + 1
        *queue = append(*queue, to)
    }
}
