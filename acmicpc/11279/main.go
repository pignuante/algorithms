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

type Heap []int

func (h Heap) Len() int {
    return len(h)
}
func (h Heap) Less(i, j int) bool {
    return h[i] > h[j]
}
func (h Heap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
    old := *h
    x := old[len(old)-1]
    new_ := old[:len(old)-1]
    *h = new_
    return x
}

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    h := Heap{}
    heap.Init(&h)
    N := nextInt()
    for i := 0; i < N; i++ {
        v := nextInt()
        if v == 0 {
            if len(h) > 0 {
                wt.WriteString(strconv.Itoa(heap.Pop(&h).(int)))
                wt.WriteByte('\n')
            } else {
                wt.WriteString("0")
                wt.WriteByte('\n')
            }
        } else {
            heap.Push(&h, v)
        }
    }
}
