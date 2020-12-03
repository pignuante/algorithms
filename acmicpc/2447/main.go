package main

import (
    "bufio"
    "os"
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
    for y := 0; y < N; y++ {
        for x := 0; x < N; x++ {
            drawStars(y, x, N)
        }
        wt.WriteString("\n")
    }
}

func drawStars(i, j, num int) {
    if (i/num)%3 == 1 && (j/num)%3 == 1 {
        wt.WriteString(" ")
    } else {
        if num/3 == 0 {
            wt.WriteString("*")
        } else {
            drawStars(i, j, num/3)
        }
    }
}
