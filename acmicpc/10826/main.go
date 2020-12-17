package main

import (
    "bufio"
    "math/big"
    "os"
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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    n := nextInt()
    r := Fibonacci(n)
    wt.WriteString(r.String())
}

func Fibonacci(num int) big.Int {
    var a, b, t big.Int
    a.SetInt64(0)
    b.SetInt64(1)
    for i := 1; i <= num; i++ {
        t.Add(&a, &b)
        a.Set(&b)
        b.Set(&t)
    }
    return a
}
