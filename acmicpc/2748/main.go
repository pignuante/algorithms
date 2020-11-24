package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())
    fibo := make([]int, n+1)
    fibo[0], fibo[1] = 0, 1
    for i := 2; i <= n; i++ {
        fibo[i] = fibo[i-2] + fibo[i-1]
    }
    wt.WriteString(strconv.Itoa(fibo[n]))
}
