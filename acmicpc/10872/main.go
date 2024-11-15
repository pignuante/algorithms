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
    N := nextInt()
    wt.WriteString(strconv.Itoa(Fatorial(N)))
}

func Fatorial(num int) int {
    if num == 0{
        return 1
    }
    return num * Fatorial(num-1)
}
