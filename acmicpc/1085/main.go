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

    x, y, w, h := nextInt(), nextInt(), nextInt(), nextInt()
    wt.WriteString(strconv.Itoa(Min(Min(w-x, x), Min(h-y, y))))
}

func Min(x, y int) (r int) {
    if x < y {
        r = x
    } else {
        r = y
    }
    return
}
