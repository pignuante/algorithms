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

    N := nextInt()
    num := 0
    for ; N > 0; N-- {
        count := 0
        n := nextInt()
        for i := 1; i <= n; i++ {
            if n%i == 0 {
                count++
                if count > 2 {
                    break
                }
            }
        }
        if count == 2 {
            num++
        }
    }
    wt.WriteString(strconv.Itoa(num))
}
