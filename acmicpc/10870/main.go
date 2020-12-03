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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    n := nextInt()
    wt.WriteString(strconv.Itoa(Fibonacci(n)))
}

func Fibonacci(num int) int {
    if num == 0 {
        return 0
    } else if num == 1 {
        return 1
    }
    return Fibonacci(num-2) + Fibonacci(num-1)
}
