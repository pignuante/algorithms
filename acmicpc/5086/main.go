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

    for {
        a, b := nextInt(), nextInt()
        if a == 0 && b == 0 {
            break
        }
        if b%a == 0 {
            wt.WriteString("factor\n")
        } else if a%b == 0 {
            wt.WriteString("multiple\n")
        } else {
            wt.WriteString("neither\n")
        }
    }
}
