package main

import (
    "bufio"
    "os"
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
    Value byte
    Next  *Node
}

type Stack struct {
    Top *Node
}

func (stack *Stack) Push(value byte) {
    node := Node{Value: value}
    if (*stack).Top != nil {
        node.Next = (*stack).Top
    }
    (*stack).Top = &node
}

func (stack *Stack) Pop() (value byte, ok bool) {
    if (*stack).Top != nil {
        value = (*stack).Top.Value
        (*stack).Top = (*stack).Top.Next
        ok = true
    }
    return
}

type Editor struct {
    Left, Right *Stack
}

func (editor *Editor) MoveLeft() {
    left, right := (*editor).Left, (*editor).Right
    if value, ok := left.Pop(); ok {
        right.Push(value)
    }
}
func (editor *Editor) MoveRight() {
    left, right := (*editor).Left, (*editor).Right
    if value, ok := right.Pop(); ok {
        left.Push(value)
    }
}
func (editor *Editor) Delete() {
    left := (*editor).Left
    _, _ = left.Pop()
}

func (editor *Editor) Write(value byte) {
    left := (*editor).Left
    left.Push(value)
}
func (editor Editor) Print() {
    for {
        if v, ok := editor.Left.Pop(); ok {
            editor.Right.Push(v)
        } else {
            break
        }
    }
    for {
        if v, ok := editor.Right.Pop(); ok {
            wt.WriteByte(v)
        } else {
            break
        }
    }
}

func InitEditor() *Editor {
    return &Editor{Left: &Stack{}, Right: &Stack{}}

}

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    editor := InitEditor()
    sc.Scan()
    for _, v := range sc.Bytes() {
        editor.Write(v)
    }
    comandNum := nextInt()
    for i := 0; i < comandNum; i++ {
        sc.Scan()
        switch sc.Bytes()[0] {
        case 'L':
            editor.MoveLeft()
        case 'D':
            editor.MoveRight()
        case 'B':
            editor.Delete()
        case 'P':
            sc.Scan()
            editor.Write(sc.Bytes()[0])
        }
    }
    editor.Print()
}
