package main

import (
    "bufio"
    "os"
    // "strconv"
)

type Parentheses struct {
    bracket rune
    before  *Parentheses
}
type Stack struct {
    head *Parentheses
    size int
}

func (stack *Stack) Push(str rune) {
    before := stack.head
    e := Parentheses{bracket: str, before: before}
    stack.head = &e
    stack.size++
}

func (stack *Stack) Pop() rune {
    if stack.size <= 0 {
        return '?'
    }
    node := stack.head
    stack.head = node.before
    stack.size--
    return node.bracket
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanLines)

    defer wt.Flush()

    end := "."
    for {
        sc.Scan()
        t := sc.Text()
        if t == end {
            break
        }
        var stack Stack
        flag := "yes"
        for _, v := range t {
            if v == '(' || v == '[' {
                stack.Push(v)
            } else if v == ')' {
                p := stack.Pop()
                if p != '(' {
                    flag = "no"
                    break
                }
            } else if v == ']' {
                p := stack.Pop()
                if p != '[' {
                    flag = "no"
                    break
                }
            }
        }
        if stack.size == 0 && flag=="yes"{
            wt.WriteString("yes" + "\n")
        } else {
            wt.WriteString("no" + "\n")
        }
    }
}
