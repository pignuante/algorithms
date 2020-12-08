package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r, _ = strconv.Atoi(sc.Text())
    return
}
func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    T := nextInt()
    for i := 0; i < T; i++ {
        x1, y1, r1 := nextInt(), nextInt(), nextInt()
        x2, y2, r2 := nextInt(), nextInt(), nextInt()
        dist := (x2-x1) * (x2-x1)+(y2-y1) * (y2-y1)
        rd := (r2-r1) * (r2-r1)
        rs := (r2+r1) * (r2+r1)

        c := "0"
        if dist == 0 && rd == 0 {
            c = "-1"
        } else {
            if dist == rd || dist == rs {
                c = "1"
            } else if dist < rs && dist > rd {
                c = "2"
            }
        }
        wt.WriteString(c + "\n")
    }
}
