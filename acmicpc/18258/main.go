package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

type Node struct {
    num  int
    next *Node
}

type Queue struct {
    head, tail *Node
    size       int
}

func (queue *Queue) Push(num int) {
    node := Node{num: num, next: nil}
    if queue.size == 0 {
        queue.head = &node
        queue.tail = &node
    } else {
        queue.tail.next = &node
        queue.tail = &node
    }
    queue.size++
}

func (queue *Queue) Pop() (num int) {
    if queue.size == 0 {
        num = -1
    } else {
        node := queue.head
        queue.head = node.next
        num = node.num
        queue.size--
    }
    return
}

func (queue Queue) Size() int {
    return queue.size
}

func (queue Queue) Empty() (isEmpty int) {
    if queue.size == 0 {
        isEmpty = 1
    }
    return isEmpty
}

func (queue Queue) Front() (num int) {
    if queue.size == 0 {
        num = -1
    } else {
        num = queue.head.num
    }
    return
}
func (queue Queue) Back() (num int) {
    if queue.size == 0 {
        num = -1
    } else {
        num = queue.tail.num
    }
    return
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    var queue Queue
    for ; N > 0; N-- {
        sc.Scan()
        command := strings.Fields(sc.Text())
        switch command[0] {
        case "push":
            num, _ := strconv.Atoi(command[1])
            queue.Push(num)
        case "pop":
            wt.WriteString(strconv.Itoa(queue.Pop()) + "\n")
        case "size":
            wt.WriteString(strconv.Itoa(queue.Size()) + "\n")
        case "empty":
            wt.WriteString(strconv.Itoa(queue.Empty()) + "\n")
        case "front":
            wt.WriteString(strconv.Itoa(queue.Front()) + "\n")
        case "back":
            wt.WriteString(strconv.Itoa(queue.Back()) + "\n")
        }
    }
}
