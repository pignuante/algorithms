package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

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
    defer wt.Flush()

    sc.Scan()
    s1 := sc.Bytes()
    sc.Scan()
    s2 := sc.Bytes()
    LCS(&s1, &s2)
}

func LCS(s1, s2 *[]byte) {
    dp := make([][]int, len(*s2)+1, len(*s2)+1)
    for k := range dp {
        dp[k] = make([]int, len(*s1)+1, len(*s1)+1)
    }
    for k1, v1 := range *s2 {
        for k2, v2 := range *s1 {
            if v1 == v2 {
                dp[k1+1][k2+1] = dp[k1][k2] + 1
            } else {
                dp[k1+1][k2+1] = Max(dp[k1][k2+1], dp[k1+1][k2])
            }
        }
    }
    wt.WriteString(strconv.Itoa(dp[len(*s2)][len(*s1)]))
    wt.WriteByte('\n')
    Display(&dp, s1, s2)
}

func Display(dp *[][]int, s1, s2 *[]byte) {
    var r string
    y, x := len((*dp)[0])-1, len(*dp)-1
    for (*dp)[x][y] != 0 {
        if (*dp)[x][y] == (*dp)[x][y-1] {
            y--
        } else if (*dp)[x][y] == (*dp)[x-1][y] {
            x--
        } else {
            r = string((*s1)[y-1]) + r
            y--
            x--
        }
    }
    wt.WriteString(r)
}
func Max(a, b int) (max int) {
    max = a
    if a < b {
        max = b
    }
    return
}
