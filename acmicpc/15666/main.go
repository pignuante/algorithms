package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var (
    N, M        int
    nums, stack []int
    visited     []bool
)

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

    N, M = nextInt(), nextInt()
    nums = make([]int, N, N)
    visited = make([]bool, N, N)
    for i := 0; i < N; i++ {
        nums[i] = nextInt()
    }
    sort.Ints(nums)
    Recursive()

}

func Recursive() {
    dup := make(map[int]bool)
    if len(stack) == M {
        for _, v := range stack {
            wt.WriteString(strconv.Itoa(v))
            wt.WriteByte(' ')
        }
        wt.WriteByte('\n')
        return
    }
    for i, v := range nums {
        if dup[v] {
            continue
        }
        if l := len(stack); l == 0 || stack[l-1] <= nums[i] {
            dup[v] = true
            stack = append(stack, v)
            Recursive()
            stack = stack[:len(stack)-1]
        }
    }
}
