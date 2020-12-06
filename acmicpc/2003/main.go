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
        r += int(c-'0')
    }
    return
}
func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, M := nextInt(), nextInt()
    A := make([]int, N+1, N+1)
    for i := 0; i < N; i++ {
        A[i] = nextInt()
    }
    wt.WriteString(strconv.Itoa(TwoPointer(A, M)))
}

func TwoPointer(A []int, M int) (count int) {
    for start, end, sum := 0, 0, A[0]; start <= end && end  <= len(A)-1; {
        if sum >= M {
            if sum == M {
                count++
            }
            sum -= A[start]
            start++
        } else {
            sum += A[end]
            end++
        }
    }
    return
}
