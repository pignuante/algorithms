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

    points := make([][2]int, 3, 3)
    for i := 0; i < 3; i++ {
        points[i] = [2]int{nextInt(), nextInt()}
    }
    x, y := 0, 0
    if points[0][0] == points[1][0] {
        x = points[2][0]
    } else if points[0][0] == points[2][0] {
        x = points[1][0]
    } else {
        x = points[0][0]
    }
    if points[0][1] == points[1][1] {
        y = points[2][1]
    } else if points[0][1] == points[2][1] {
        y = points[1][1]
    } else {
        y = points[0][1]
    }
    wt.WriteString(strconv.Itoa(x) + " " + strconv.Itoa(y))
}
