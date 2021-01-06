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
func (h *Points) Push(point interface{}) {
    *h = append(*h, point.(Point))
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
    Parents   []int
)

type Stack struct {
    Path []int
    Size int
}

func (stack *Stack) Push(value int) {
    (*stack).Path = append((*stack).Path, value)
    (*stack).Size++
}

func (stack *Stack) Pop() (value int, ok bool) {
    if (*stack).Size > 0 {
        l := len((*stack).Path) - 1
        value = (*stack).Path[l]
        (*stack).Path = (*stack).Path[:l]
        ok = true
        (*stack).Size--
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, K = nextInt(), nextInt()
    max = Max(Max(N*2+1, K*2+1), 10)
    Visited = make([]int, max+1, max+1)
    Parents = make([]int, max+1, max+1)

    time := FindBrother()
    wt.WriteString(strconv.Itoa(time))
    wt.WriteByte('\n')
    stack := Stack{}
    idx := K
    stack.Push(K)
    for i := 0; i < 5; i++ {
        if Parents[idx] == -1 {
            break
        }
        stack.Push(Parents[idx])
        idx = Parents[idx]
    }
    for range stack.Path {
        if v, ok := stack.Pop(); ok {
            wt.WriteString(strconv.Itoa(v))
            wt.WriteByte(' ')
        }
    }
}

func FindBrother() (time int) {
    for i := range Visited {
        Visited[i] = max + 1
        Parents[i] = max + 1
    }
    heap.Init(&Queue)
    heap.Push(&Queue, Point{Second: 0, Location: N})
    Visited[N], Parents[N] = 0, -1
    time = int(^uint(0) >> 1)

    idx := 0
    for len(Queue) > 0 {
        idx++
        point := heap.Pop(&Queue).(Point)
        if point.Second > time {
            continue
        }
        if point.Location == K {
            time = point.Second
            break
        }

        if loc := point.Location * 2; loc <= max && Visited[loc] > point.Second {
            Visited[loc] = point.Second
            Parents[loc] = point.Location
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: loc})
        }
        if loc := point.Location + 1; loc <= max && Visited[loc] > point.Second {
            Visited[loc] = point.Second + 1
            Parents[loc] = point.Location
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: loc})
        }
        if loc := point.Location - 1; 0 <= loc && loc <= max && Visited[loc] > point.Second {
            Visited[loc] = point.Second + 1
            Parents[loc] = point.Location
            heap.Push(&Queue, Point{Second: point.Second + 1, Location: loc})
        }
    }
    return
}

func Max(a, b int) (max int) {
    max = a
    if a < b {
        max = b
    }
    return
}
