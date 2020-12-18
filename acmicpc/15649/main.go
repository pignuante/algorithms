package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

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
    defer wt.Flush()

    N, M := nextInt(), nextInt()
    var stack []int
    visited := make([]bool, N, N)
    DFS(&stack, &visited, N, M)
}

func DFS(stack *[]int, visited *[]bool, n, m int) {
    if len(*stack) == m { // 조건에 부합하면 (종료조건)
        for _, v := range *stack {
            wt.WriteString(strconv.Itoa(v))
            wt.WriteByte(' ')
        }
        wt.WriteByte('\n')
        return
    }
    for i := 0; i < n; i++ {
        if (*visited)[i] == false { // 전진 조건
            (*visited)[i] = true
            (*stack) = append((*stack), i+1)
            DFS(stack, visited, n, m)
            (*stack) = (*stack)[:len(*stack)-1]
            (*visited)[i] = false
        }
    }
}
