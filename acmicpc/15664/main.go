package main

import (
    "bufio"
    "os"
    "sort"
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

var (
    N, M    int
    stack   []int
    nums    []int
    visited []bool
)

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, M = nextInt(), nextInt()
    nums, visited = make([]int, N, N), make([]bool, N, N)
    for i := 0; i < N; i++ {
        nums[i] = nextInt()
    }
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    Recursive()
}

func Recursive() {
    dup := make(map[int]bool, N)
    if len(stack) == M {
        for _, v := range stack {
            wt.WriteString(strconv.Itoa(v))
            wt.WriteByte(' ')
        }
        wt.WriteByte('\n')
        return
    }
    for i, v := range nums {
        if dup[v] || visited[i] {
            continue
        }
        if l := len(stack); l == 0 || stack[l-1] <= nums[i] {
            dup[v] = true
            visited[i] = true
            (stack) = append(stack, v)
            Recursive()
            stack = stack[:len(stack)-1]
            visited[i] = false
        }
    }
}
