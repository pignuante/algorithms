package main

import (
    "bufio"
    "container/heap"
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

type Point struct {
    Second, Location int
}
type Points []Point

func (h Points) Len() int {
    return len(h)
}
func (h Points) Less(i, j int) bool {
    return h[i].Second < h[j].Second
}
func (h Points) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *Points) Push(element interface{}) {
    *h = append(*h, element.(Point))
}
func (h *Points) Pop() interface{} {
    old := *h
    n := len(old)
    element := old[n-1]
    *h = old[0 : n-1]
    return element
}

var (
    N, K    int
    Queue   Points
    Visited []int
)
const MAX = 100000

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, K = nextInt(), nextInt()
    Visited = make([]int, MAX+1, MAX+1)
    FindBrother()
}

func FindBrother() {
    for i := range Visited {
        Visited[i] = MAX+1
    }
    heap.Init(&Queue)
    heap.Push(&Queue, Point{Second: 0, Location: N})
    Visited[N] = 0
    for len(Queue) > 0 {
        point := heap.Pop(&Queue).(Point)

        if point.Location == K {
            wt.WriteString(strconv.Itoa(point.Second))
            break
        }
        if l := point.Location * 2; l <= MAX && Visited[l] > point.Second {
            Visited[l] = point.Second
            heap.Push(&Queue, Point{Second: point.Second, Location: l})
        }
        if l := point.Location + 1; l <= MAX && Visited[l] > point.Second {
            Visited[l] = point.Second + 1
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: l})
        }
        if l := point.Location - 1; 0 <= l && l <= MAX && Visited[l] > point.Second {
            Visited[l] = point.Second + 1
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: l})
        }
    }
}
