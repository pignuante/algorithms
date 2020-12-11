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

    girl, boy, intern := nextInt(), nextInt(), nextInt()
    wt.WriteString(strconv.Itoa(Min(Min(girl/2, boy), (girl+boy-intern)/3)))

}

func Min(a, b int) (min int) {
    min = a
    if a > b {
        min = b
    }
    return
}
