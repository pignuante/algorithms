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

    E, S, M := nextInt(), nextInt(), nextInt()

    year := 0
    for {
        year++
        if (year-E)%15 == 0 && (year-S)%28 == 0 && (year-M)%19 == 0 {
            wt.WriteString(strconv.Itoa(year))
            break
        }
    }
}
