package main

import (
    "bufio"
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

type Node struct {
    Value        int
    Before, Next *Node
}

type Ring struct {
    Head, Tail  *Node
    Size, Count int
}

func (ring *Ring) PushBack(value int) {
    node := Node{Value: value}
    if (*ring).Size == 0 {
        node.Before = (*ring).Tail
        (*ring).Tail = &node
        (*ring).Head = (*ring).Tail
    } else {
        node.Before = (*ring).Tail
        (*ring).Tail.Next = &node
        (*ring).Tail = &node
    }
    (*ring).Size++
}

func (ring *Ring) PushFront(value int) {
    node := Node{Value: value}
    if (*ring).Size == 0 {
        (*ring).Head = &node
        (*ring).Tail = (*ring).Head
    } else {
        head := (*ring).Head
        node.Next = head
        (*ring).Head = &node
    }
    (*ring).Size++
}

func (ring *Ring) PopFront() (value int, ok bool) {
    node := (*ring).Head
    if (*ring).Size > 1 {
        (*ring).Head = node.Next
        value, ok = node.Value, true
        (*ring).Size--
    } else if (*ring).Size == 1 {
        (*ring).Head = nil
        (*ring).Tail = nil
        value, ok = node.Value, true
        (*ring).Size--
    }
    return
}

func (ring *Ring) PopBack() (value int, ok bool) {
    node := (*ring).Tail
    if (*ring).Size > 1 {
        (*ring).Tail = node.Before
        (*ring).Size--
        value, ok = node.Value, true
    } else {
        (*ring).Head = nil
        (*ring).Tail = nil
        (*ring).Size--
        value, ok = node.Value, true
    }
    return
}

func (ring *Ring) MoveLeft() {
    if (*ring).Size > 0 {
        v, ok := (*ring).PopFront()
        if ok {
            (*ring).PushBack(v)
        }
        (*ring).Count++
    }
}

func (ring *Ring) MoveRight() {
    if (*ring).Size > 0 {
        v, ok := (*ring).PopBack()
        if ok {
            (*ring).PushFront(v)
        }
        (*ring).Count++
    }
}
func (ring *Ring) GetCount(value int) (count int) {
    l := 0
    for ring.Head.Value != value {
        ring.MoveLeft()
        l++
    }
    count = Min(l, ring.Size-l)
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, M := nextInt(), nextInt()
    nums := make([]int, M, M)
    var ring Ring
    for i := 1; i <= N; i++ {
        ring.PushBack(i)
    }
    for i := 0; i < M; i++ {
        nums[i] = nextInt()
    }
    result := 0
    for _, v := range nums {
        result += ring.GetCount(v)
        ring.PopFront()
    }
    wt.WriteString(strconv.Itoa(result))
}

func Min(a, b int) (min int) {
    min = a
    if a > b {
        min = b
    }
    return
}
