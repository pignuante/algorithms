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

    max, sum := 0, 0
    for i := 0; i < 4; i++ {
        sum += -nextInt() + nextInt()
        if sum > max {
            max = sum
        }
    }
    wt.WriteString(strconv.Itoa(max))
}
