package main

import (
    "bufio"
    "os"
    "strconv"
)

type Element struct {
    value int
    next  *Element
}
type Stack struct {
    head *Element
    sum  int
}

func (stack *Stack) Push(num int) {
    head := stack.head
    node := Element{value: num, next: head}
    stack.head = &node
    stack.sum += num
}
func (stack *Stack) Pop() {
    head := stack.head
    stack.head = head.next
    stack.sum -= head.value
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()

    sc.Scan()
    K, _ := strconv.Atoi(sc.Text())
    var stack Stack
    for ; K > 0; K-- {
        sc.Scan()
        n, _ := strconv.Atoi(sc.Text())
        if n == 0 {
            stack.Pop()
        } else {
            stack.Push(n)
        }
    }
    wt.WriteString(strconv.Itoa(stack.sum))
}
