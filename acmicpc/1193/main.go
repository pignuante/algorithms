package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r = 0
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
    X := nextInt()
    r := findFraction(X)
    wt.WriteString(strconv.Itoa(r[0]) + "/" + strconv.Itoa(r[1]))
}

func findFraction(X int) [2]int {
    i, s := 2, 1

    for X > s {
        s += i
        i += 1
    }
    a := [2]int{s - X + 1, i - s + X - 1}[i%2]
    return [2]int{a, i - a}
}
