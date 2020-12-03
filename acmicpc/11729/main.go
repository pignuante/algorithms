package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

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

    defer wt.Flush()
    N := nextInt()
    var moved [][2]int
    Hanoi(N, 1, 2, 3, &moved)
    wt.WriteString(strconv.Itoa(len(moved)) + "\n")
    for _, v := range moved {
        wt.WriteString(strconv.Itoa(v[0]) + " " + strconv.Itoa(v[1]) + "\n")
    }
}

func Hanoi(num, start, mid, end int, moved *[][2]int) {
    if num == 1 {
        *moved = append(*moved, [2]int{start, end})
    } else {
        Hanoi(num-1, start, end, mid, moved)
        *moved = append(*moved, [2]int{start, end})
        Hanoi(num-1, mid, start, end, moved)
    }
    return
}
