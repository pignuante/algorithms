package main

import (
    "bufio"
    "os"
    "sort"
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

    N, C := nextInt(), nextInt()
    routers := make([]int, N, N)
    for i := 0; i < N; i++ {
        routers[i] = nextInt()
    }
    sort.Ints(routers)

    start, end, result := 1, routers[N-1], 0
    for start <= end {
        mid := (start + end) >> 1
        count := countRouter(routers, mid)
        if count >= C {
            result = mid
            start = mid + 1
        } else if count < C {
            end = mid - 1
        }
    }
    wt.WriteString(strconv.Itoa(result))
}

func countRouter(routers []int, dist int) (count int) {
    curDist, count := routers[0], 1
    for _, router := range routers[1:] {
        if curDist+dist <= router {
            count++
            curDist = router
        }
    }
    return
}
