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

    for {
        k := nextInt()
        if k == 0 {
            break
        }
        nums := make([]int, k, k)
        visited := make([]bool, k, k)
        for i := 0; i < k; i++ {
            nums[i] = nextInt()
        }
        var stack []int
        DFS(&nums, &stack, &visited)
        wt.WriteByte('\n')
    }
}

func DFS(nums *[]int, stack *[]int, visited *[]bool) {
    if len(*stack) == 6 {
        for _, v := range *stack {
            wt.WriteString(strconv.Itoa(v))
            wt.WriteByte(' ')
        }
        wt.WriteByte('\n')
        return
    }
    for i, v := range *nums {
        if (*visited)[i] {
            continue
        }
        if l := len(*stack); l == 0 || (*stack)[l-1] < v {
            (*visited)[i] = true
            *stack = append((*stack), v)
            DFS(nums, stack, visited)
            *stack = (*stack)[:len(*stack)-1]
            (*visited)[i] = false
        }
    }
}
