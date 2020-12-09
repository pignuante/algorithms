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

    N, count := nextInt(), 0
    for i := 0; i < 5; i++ {
        if car := nextInt(); car == N {
            count++
        }
    }
    wt.WriteString(strconv.Itoa(count))
}
