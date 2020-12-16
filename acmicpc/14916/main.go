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

    n := nextInt()

    f, r := Changes(n)
    if f == 0 {
        wt.WriteString(strconv.Itoa(r))
    } else {
        wt.WriteString("-1")
    }
}

func Changes(n int) (int, int) {
    r := 0
    for n > 0 {
        if n%5 == 0 {
            r = r + (n / 5)
            n = 0
        } else {
            n -= 2
            r++
        }
    }

    return n, r
}
