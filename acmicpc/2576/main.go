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
    sum, min := 0, 101

    for i := 0; i < 7; i++ {
        num := nextInt()
        if num%2 == 1 {
            sum += num
            if num < min {
                min = num
            }
        }
    }
    if min == 101 {
        wt.WriteString("-1")
    } else {
        wt.WriteString(strconv.Itoa(sum))
        wt.WriteByte('\n')
        wt.WriteString(strconv.Itoa(min))
    }
}
