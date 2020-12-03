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
    N, k := nextInt(), nextInt()
    r := BinarySearch(N, k)
    wt.WriteString(strconv.Itoa(r) + "\n")
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func BinarySearch(N, k int) (count int) {
    for start, end := 1, k; start <= end; {
        sum, mid := 0, (start+end)>>1
        for i := 1; i <= N; i++ {
            sum += min(mid/i, N)
        }
        if sum < k {
            start = mid + 1
        } else {
            end = mid - 1
            count = mid
        }
    }
    return
}
