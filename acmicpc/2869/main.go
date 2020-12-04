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
    A, B, V := nextInt(), nextInt(), nextInt()
    wt.WriteString(strconv.Itoa(snail(A, B, V)))
}

func snail(A, B, V int) (days int) {
    days = (V-B-1)/(A-B) + 1
    return
}
