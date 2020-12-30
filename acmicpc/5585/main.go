package main

import (
    "bufio"
    "os"
    "strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

const Money = 1000

var coins = [...]int{500, 100, 50, 10, 5, 1}

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    m := nextInt()
    change := Money - m
    count := 0
    for _, coin := range coins {
        for change >= coin {
            change -= coin
            count++
        }
    }
    wt.WriteString(strconv.Itoa(count))
}
