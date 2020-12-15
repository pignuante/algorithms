package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    T, _ := strconv.Atoi(sc.Text())
    for ; T > 0; T-- {
        s := 0
        sc.Scan()
        for _, v := range sc.Bytes() {
            if v == ',' {
                continue
            }
            s += int(v - '0')
        }
        wt.WriteString(strconv.Itoa(s))
        wt.WriteByte('\n')
    }
}
