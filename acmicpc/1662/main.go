package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

type result struct {
    count, idx int
}

var S []byte

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    sc.Scan()
    S = sc.Bytes()

    wt.WriteString(strconv.Itoa(Recursive()))
}

func Recursive() int {
    r := 0
    for {
        switch {
        case len(S) == 0:
            return r
        case S[0] == ')':
            S = S[1:]
            return r
        case len(S) > 1 && S[1] == '(':
            k := int(S[0]) - '0'
            S = S[2:]
            a := Recursive()
            r += k * a
        default:
            S = S[1:]
            r++
        }
    }
}
