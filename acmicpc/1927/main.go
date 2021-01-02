package main

import (
    "bufio"
    "container/heap"
    "os"
    "strconv"
)

type Heap []int

func (heap Heap) Len() int {
    return len(heap)
}

func (heap Heap) Less(i, j int) bool {
    return heap[i] < heap[j]
}

func (heap Heap) Swap(i, j int) {
    heap[i], heap[j] = heap[j], heap[i]
}

func (nums *Heap) Push(element interface{}) {
    *nums = append(*nums, element.(int))
}
func (nums *Heap) Pop() interface{} {
    n := len(*nums)
    element := (*nums)[n-1]
    *nums = (*nums)[:n-1]
    return element
}

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

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    h := &Heap{}
    heap.Init(h)

    N := nextInt()
    for i := 0; i < N; i++ {
        switch t := nextInt(); t {
        case 0:
            if len(*h) > 0 {
                v := heap.Pop(h)
                wt.WriteString(strconv.Itoa(v.(int)))
            } else {
                wt.WriteString("0")
            }
            wt.WriteByte('\n')
        default:
            heap.Push(h, t)
        }
    }
}
