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
        r += int(c-'0')
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    a, b := nextInt(), nextInt()
    gcd := GCD(a, b)
    lcm := LCM(a, b, gcd)
    wt.WriteString(strconv.Itoa(gcd))
    wt.WriteByte('\n')
    wt.WriteString(strconv.Itoa(lcm))
}

func GCD(a, b int) (int) {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func LCM (a, b, r int) (int) {
    return a*b / r
}
