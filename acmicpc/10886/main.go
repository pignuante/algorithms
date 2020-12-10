package main

import (
    "bufio"
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
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    sum := 0
    for i := 0; i < N; i++ {
        sum += nextInt()
    }
    if sum > N/2 {
        wt.WriteString("Junhee is cute!")
    } else {
        wt.WriteString("Junhee is not cute!")
    }
}
