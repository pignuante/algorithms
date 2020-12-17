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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    n := nextInt()
    wt.WriteString(strconv.Itoa(Fibonacci(n)))
}

func Fibonacci(num int) int {
    a, b := 0, 1
    for i := 1; i <= num; i++ {
        a, b = b, (a+b)%1000000007
    }
    return a
}
