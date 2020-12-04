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

    wt.WriteString(strconv.Itoa(sequence(nextInt())))
}

func sequence(num int) (moved int) {
    moved = 1

    for i := 1; num > 1; i++ {
        num = num - 6*i
        moved++
    }
    return
}
