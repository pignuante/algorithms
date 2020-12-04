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

    T := nextInt()
    for i := 0; i < T; i++ {
        room := roomAssign(nextInt(), nextInt(), nextInt())
        wt.WriteString(room)
    }
}

func roomAssign(H, W, n int) (room string) {
    w := (n-1)/H
    h := n-w*H
    return strconv.Itoa(h*100+w+1)+"\n"
}
