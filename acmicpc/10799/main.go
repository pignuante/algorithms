package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    sc.Buffer(make([]byte, 110000), 110000)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    sticks := sc.Bytes()
    r, k := 0, 0
    for i, c := range sticks {
        switch c {
        case '(':
            k++
        case ')':
            k--
            if sticks[i-1] == '(' {
                r += k
            } else {
                r++
            }
        }
    }
    wt.WriteString(strconv.Itoa(r))
}
