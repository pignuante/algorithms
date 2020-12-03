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
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    A, B, C := nextInt(), nextInt(), nextInt()
    wt.WriteString(strconv.Itoa(breakEvenPoint(A, B, C)))
}

func breakEvenPoint(fixed, product, sell int) (point int) {
    if product >= sell {
        point = -1
    } else {
        point = (fixed / (sell - product)) + 1
    }
    return
}
