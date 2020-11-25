package main

import (
    "bufio"
    "os"
    "strconv"
)

type Element struct {
    Num  int
    Next *Element
}
type Stack struct {
    head *Element
    size int
}

func (stack *Stack) Push(X int) {
    e := Element{Num: X, Next: stack.head}
    stack.head = &e
    stack.size++
}
func (stack *Stack) Pop() int {
    if stack.size == 0 {
        return -1
    }
    e := stack.head
    stack.head = e.Next
    stack.size--
    return e.Num
}
func (stack Stack) Size() int {
    return stack.size
}
func (stack Stack) Empty() int {
    if stack.size == 0 {
        return 1
    }
    return 0
}
func (stack Stack) Top() int {
    if stack.size == 0 {
        return -1
    }
    return stack.head.Num
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanWords)

    defer wt.Flush()

    var stack Stack

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    for ; N > 0; N-- {
        sc.Scan()
        text := sc.Text()
        switch text {
        case "push":
            sc.Scan()
            v, _ := strconv.Atoi(sc.Text())
            stack.Push(v)
        case "pop":
            v := stack.Pop()
            wt.WriteString(strconv.Itoa(v) + "\n")
        case "size":
            s := stack.Size()
            wt.WriteString(strconv.Itoa(s) + "\n")
        case "empty":
            v := stack.Empty()
            wt.WriteString(strconv.Itoa(v) + "\n")
        case "top":
            v := stack.Top()
            wt.WriteString(strconv.Itoa(v) + "\n")
        }
    }
}
