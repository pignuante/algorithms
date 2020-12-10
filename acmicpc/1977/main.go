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

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    M, N := nextInt(), nextInt()
    sum, min := 0, 100000001
    for x := 1; x < N; x++ {
        if t := x * x; M <= t && t <= N {
            sum += t
            if min > t {
                min = t
            }
        } else if t > N {
            break
        }
    }
    if min == 100000001 {
        wt.WriteString("-1")
    } else {
        wt.WriteString(strconv.Itoa(sum))
        wt.WriteByte('\n')
        wt.WriteString(strconv.Itoa(min))
    }
}
