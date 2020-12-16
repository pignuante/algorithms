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
        r += int(c-'0')
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    A, B, C, N := nextInt(), nextInt(), nextInt(), nextInt()
    var dp [301]int
    dp[0] = 1
    for i := 1 ; i <= N; i++ {
        if i >= A {
            if i >= B  {
                if i >= C {
                    dp[i] = dp[i-A] + dp[i-B] + dp[i-C]
                } else {
                    dp[i] = dp[i - A] + dp[i - B]
                }
            } else {
                dp[i] = dp[i-A]
            }
        } else {
            dp[i] = 0
        }
    }
    if dp[N] != 0 {
        wt.WriteString("1")
    } else {
        wt.WriteString("0")
    }
}
