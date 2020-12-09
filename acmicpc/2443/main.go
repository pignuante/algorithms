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

    N := nextInt()

    for i := 0; i < N; i++ {
        str := ""
        for space := 0; space < i; space++ {
            str += " "
        }
        for star := 0; star < 2*N-i*2-1; star++ {
            str += "*"
        }
        wt.WriteString(str + "\n")
    }
}
