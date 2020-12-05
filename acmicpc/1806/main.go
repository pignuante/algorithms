package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r = 0
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}
func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, S := nextInt(), nextInt()
    LIS := make([]int, N+1, N+1)
    for i := 0; i < N; i++ {
        LIS[i] = nextInt()
    }
    wt.WriteString(strconv.Itoa((DoublePointer(LIS, S))))
}

func DoublePointer(LIS []int, S int) (l int) {
    l = len(LIS) + 1
    start, end, sum := 0, 0, LIS[0]
    for start <= end && end < len(LIS)-1 {
        if sum >= S {
            if t := end - start + 1; l > t {
                l = t
            }
            sum -= LIS[start]
            start++
        } else if sum < S {
            end += 1
            sum += LIS[end]
        }
    }
    if l == len(LIS)+1 {
        l = 0
    }
    return
}
