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

    M, N := nextInt(), nextInt()
    Sieve := Eratosthenes(N)
    for i := M; i <= N; i++ {
        if !Sieve[i] {
            wt.WriteString(strconv.Itoa(i) + "\n")
        }
    }
}

func Eratosthenes(n int) (Sieve []bool) {
    Sieve = make([]bool, n+1, n+1)
    Sieve[0], Sieve[1] = true, true

    for i := 2; i <= n; i++ {
        if Sieve[i] {
            continue
        }
        for j := i + i; j <= n; j += i {
            Sieve[j] = true
        }
    }

    return
}
