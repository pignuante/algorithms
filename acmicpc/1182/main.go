package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    flag := 1
    for _, c := range sc.Bytes() {
        if c == '-' {
            flag = -1
            continue
        }
        r *= 10
        r += int(c - '0')
    }
    r *= flag
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, S := nextInt(), nextInt()
    nums := make([]int, N, N)
    for i := 0; i < N; i++ {
        nums[i] = nextInt()
    }
    count := 0
    DFS(N, S, 0, 0, &count, &nums)
    wt.WriteString(strconv.Itoa(count))
}

func DFS(N, S, idx, value int, count *int, nums *[]int) {
    if value+(*nums)[idx] == S {
        *count++
    }
    if idx+1 == N {
        return
    }
    DFS(N, S, idx+1, value+(*nums)[idx], count, nums)
    DFS(N, S, idx+1, value, count, nums)
}
