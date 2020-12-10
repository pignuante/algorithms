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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    names := make([][]byte, N, N)
    for i := 0; i < N; i++ {
        sc.Scan()
        names[i] = sc.Bytes()
    }
    str := names[0]
    for _, v := range names[1:] {
        for kk := range v {
            if str[kk] != v[kk] {
                str[kk] = byte('?')
            }
        }
    }
    wt.WriteString(string(str))
}
