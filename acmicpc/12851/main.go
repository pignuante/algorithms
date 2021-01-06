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
    N, K, max int
    Queue     Points
    Visited   []int
)

const MAX = 100000

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, K = nextInt(), nextInt()
    max = Max(Max(N*2+1, K*2+1), MAX)
    Visited = make([]int, max+1, max+1)
    time, count := FindBrother()
    wt.WriteString(strconv.Itoa(time))
    wt.WriteByte('\n')
    wt.WriteString(strconv.Itoa(count))
}

func FindBrother() (time, count int) {
    for i := range Visited {
        Visited[i] = MAX + 1
    }
    heap.Init(&Queue)
    heap.Push(&Queue, Point{Second: 0, Location: N})
    Visited[N] = 0
    count, time = 0, int(^uint(0)>>1)
    idx := 0
    for len(Queue) > 0 {
        point := heap.Pop(&Queue).(Point)
        idx++
        if point.Second > time {
            continue
        }
        if point.Location == K {
            if time > point.Second {
                time = point.Second
                count = 1
            } else {
                count++
            }
        }
        if l := point.Location * 2; l <= max && Visited[l] >= point.Second {
            Visited[l] = point.Second
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: l})
        }
        if l := point.Location + 1; l <= max && Visited[l] >= point.Second {
            Visited[l] = point.Second + 1
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: l})
        }
        if l := point.Location - 1; 0 <= l && l <= max && Visited[l] >= point.Second {
            Visited[l] = point.Second + 1
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: l})
        }
    }
    return time, count
}

func Max(a, b int) (m int) {
    m = a
    if a < b {
        m = b
    }
    return
}
