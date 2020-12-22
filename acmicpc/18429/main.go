package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

const weight = 500

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, K := nextInt(), nextInt()
    kit, visited := make([]int, N, N), make([]bool, N, N)
    var stack []int
    count := 0
    for i := 0; i < N; i++ {
        kit[i] = nextInt()
    }
    DFS(&kit, &stack, &visited, weight, K, &count)
    wt.WriteString(strconv.Itoa(count))
}

func DFS(kit, stack *[]int, visited *[]bool, value, K int, count *int) {
    if value < weight {
        return
    }
    if len(*stack) == len(*kit) {
        *count++
        return
    }
    for i := 0; i < len(*kit); i++ {
        if (*visited)[i] {
            continue
        }
        (*visited)[i] = true
        (*stack) = append(*stack, (*kit)[i])
        DFS(kit, stack, visited, value+(*kit)[i]-K, K, count)
        (*stack) = (*stack)[:len(*stack)-1]
        (*visited)[i] = false
    }
}
