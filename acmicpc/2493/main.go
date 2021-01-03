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
    Height, Idx int
    Next        *Node
}

type Stack struct {
    Top  *Node
    Size int
}

func (stack *Stack) Push(node Node) {
    if (*stack).Top != nil {
        top := (*stack).Top
        node.Next = top
    }
    (*stack).Top = &node
    (*stack).Size++
}

func (stack *Stack) Pop() {
    if (*stack).Size > 0 {
        (*stack).Top = (*stack).Top.Next
    }
    stack.Size--
}

func (stack Stack) Peek() (node Node) {
    if stack.Size > 0 {
        node = *stack.Top
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N := nextInt()
    stack := Stack{}
    result := make([]int, N, N)
    for i := 0; i < N; i++ {
        height := nextInt()
        for stack.Size > 0 {
            if peek := stack.Peek(); peek.Height >= height {
                result[i] = peek.Idx
                break
            }
            stack.Pop()
        }
        if stack.Size == 0 {
            result[i] = 0
        }
        stack.Push(Node{Height: height, Idx: i + 1})
    }

    for _, v := range result {
        wt.WriteString(strconv.Itoa(v))
        wt.WriteByte(' ')
    }
}
