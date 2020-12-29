package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
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

    L, P := nextInt(), nextInt()
    for i := 0; i < 5; i++ {
        wt.WriteString(strconv.Itoa(nextInt() - L*P))
        wt.WriteByte(' ')
    }
}
