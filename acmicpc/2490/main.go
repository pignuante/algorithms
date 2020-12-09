package main

import (
    "bufio"
    "os"
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
    for i := 0; i < 3; i++ {
        s := 0
        for j := 0; j < 4; j++ {
            s += nextInt()
        }
        switch s {
        case 4:
            wt.WriteByte('E')
        case 3:
            wt.WriteByte('A')
        case 2:
            wt.WriteByte('B')
        case 1:
            wt.WriteByte('C')
        case 0:
            wt.WriteByte('D')
        }
        wt.WriteByte('\n')
    }
}
