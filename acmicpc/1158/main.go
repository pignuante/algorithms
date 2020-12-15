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

type Node struct {
    value int
    next  *Node
}

type Queue struct {
    head, tail *Node
}

func (queue *Queue) Push(value int) {
    node := &Node{value: value, next: nil}
    if queue.head == nil {
        queue.head = node
        queue.tail = node
    } else {
        tail := queue.tail
        tail.next = node
        queue.tail = node
    }
    node.next = queue.head
}

func (queue *Queue) Pop(k int) (r int) {
    for i := 0; i < k-1; i++ {
        queue.head = queue.head.next
        queue.tail = queue.tail.next
    }
    node := queue.head
    r = node.value
    queue.tail.next = node.next
    queue.head = node.next
    node.next = nil
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, K := nextInt(), nextInt()
    var queue Queue
    for i := 1; i <= N; i++ {
        queue.Push(i)
    }
    wt.WriteByte('<')
    for i := 0; i < N; i++ {
        wt.WriteString(strconv.Itoa(queue.Pop(K)))
        if i == N-1 {
            break
        }
        wt.WriteByte(',')
        wt.WriteByte(' ')
    }
    wt.WriteByte('>')
}
