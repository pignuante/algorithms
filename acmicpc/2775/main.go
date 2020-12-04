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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    T := nextInt()
    for ; T > 0; T-- {
        k, n := nextInt(), nextInt()
        houses := make([]int, n, n)
        for i := 0; i < n; i++ {
            houses[i] = i + 1
        }
        for y := 0; y < k; y++ {
            for x := 1; x < n; x++ {
                houses[x] += houses[x-1]
            }
        }
        wt.WriteString(strconv.Itoa(houses[n-1]) + "\n")
    }
}
