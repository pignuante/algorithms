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

    T := nextInt()
    for ; T > 0; T-- {
        a, b := nextInt(), nextInt()
        res := 1
        for j := 0; j < b; j++ {
            res = (a * res) % 10
        }
        if res == 0 {
            res = 10
        }

        wt.WriteString(strconv.Itoa(res))
        wt.WriteByte('\n')
    }
}
