package main

import (
    "bufio"
    "os"
    "sort"
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
    nums := make([]int, N, N)
    for i := 0; i < N; i++ {
        nums[i] = nextInt()
    }
    sort.Ints(nums)
    var stack []int
    DFS(&nums, &stack, N, M)
}

func DFS(nums *[]int, stack *[]int, n, m int) {
    if len(*stack) == m { // 조건에 부합하면 (종료조건)
        for _, v := range *stack {
            wt.WriteString(strconv.Itoa(v))
            wt.WriteByte(' ')
        }
        wt.WriteByte('\n')
        return
    }
    for _, v := range *nums {
            (*stack) = append((*stack), v)
            DFS(nums, stack, n, m)
            (*stack) = (*stack)[:len(*stack)-1]
    }
}
