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

    N := nextInt()
    var queue Queue
    for ; N > 0; N-- {
        sc.Scan()
        cmd := sc.Text()
        switch cmd {
        case "push":
            sc.Scan()
            value := sc.Text()
            queue.Push(value)
        case "pop":
            wt.WriteString(queue.Pop())
            wt.WriteByte('\n')
        case "size":
            wt.WriteString(strconv.Itoa(queue.Size()))
            wt.WriteByte('\n')
        case "empty":
            wt.WriteString(queue.Empty())
            wt.WriteByte('\n')
        case "front":
            wt.WriteString(queue.Front())
            wt.WriteByte('\n')
        case "back":
            wt.WriteString(queue.Back())
            wt.WriteByte('\n')
        }
    }
}

type Node struct {
    value string
    next  *Node
}

type Queue struct {
    head, tail *Node
    size       int
}

func (queue *Queue) Push(x string) {
    node := &Node{value: x, next: nil}
    if queue.head == nil {
        queue.head = node
        queue.tail = node
    } else {
        q := queue.tail
        q.next = node
        queue.tail = node
    }
    queue.size++
}

func (queue *Queue) Pop() (r string) {
    if queue.size == 0 {
        r = "-1"
    } else {
        q := queue.head
        queue.head = q.next
        r = q.value
        queue.size--
    }
    return
}

func (queue Queue) Size() int {
    return queue.size
}

func (queue Queue) Empty() string {
    if queue.size == 0 {
        return "1"
    }
    return "0"
}

func (queue Queue) Front() string {
    if queue.size == 0 {
        return "-1"
    }
    return queue.head.value
}

func (queue Queue) Back() string {
    if queue.size == 0 {
        return "-1"
    }
    return queue.tail.value
}
