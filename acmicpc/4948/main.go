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

    max := -1
    var nums []int
    for n := -1; ; {
        n = nextInt()
        if n == 0 {
            break
        }
        nums = append(nums, n)
        if max < n {
            max = n
        }
    }
    Sieve := Eratosthenes(max + max)
    for _, v := range nums {
        count := 0
        for a := v + 1; a <= v+v; a++ {
            if !Sieve[a] {
                count++
            }
        }
        wt.WriteString(strconv.Itoa(count) + "\n")
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
