package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func Abs(number int) int {
	if number < 0 {
		number *= -1
	}
	return number
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	ah, bh := Abs(h[i]), Abs(h[j])
	if ah == bh {
		return h[i] < h[j]
	}
	return ah < bh
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(element interface{}) {
	*h = append(*h, element.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	N := nextInt()
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for ; N > 0; N-- {
		num := nextInt()
		if num == 0 {
			if minHeap.Len() > 0 {
				v := heap.Pop(minHeap)
				wt.WriteString(strconv.Itoa(v.(int)))
				wt.WriteByte('\n')
			} else {
				wt.WriteByte('0')
				wt.WriteByte('\n')
			}
			continue
		}
		heap.Push(minHeap, num)
	}
}
