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
    r = 0
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}
func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    T := nextInt()
    Sieve := Eratosthenes(10000)
    for ; T > 0; T-- {
        n := nextInt()
        FindGoldbach(&Sieve, n)
    }
}

func FindGoldbach(Sieve *[]bool, num int) {
    for start := num >> 1; start >= 2; start-- {
        end := num - start
        if !(*Sieve)[start] && !(*Sieve)[end] {
            wt.WriteString(strconv.Itoa(start)+ " " + strconv.Itoa(end)+"\n")
            break
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
