package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
func nextInt() (r int) {
    sc.Scan()
    r, sign := 0, 1

    for _, c := range sc.Bytes() {
        if c == '-' {
            sign = -1
            continue
        }
        r *= 10
        r += int(c-'0')
    }
    r *= sign
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    wt.WriteString(strconv.Itoa(nextInt()))
}
