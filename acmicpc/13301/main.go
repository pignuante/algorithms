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

    N := nextInt()
    wt.WriteString(strconv.Itoa(GetNumOfTile(N)))
}

func GetNumOfTile(K int) int {
    a, b := 4, 6
    for i := 1; i < K; i++ {
        a, b = b, a+b
    }
    return a
}
