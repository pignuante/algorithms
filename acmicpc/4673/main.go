package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    result := make([]bool, 10001)
    for i := 1; i < 10001; i++ {
        D := d(i)
        if D < 10001 {
            result[D] = true

        }
    }
    result[0] = true
    for k := range result {
        if !result[k] {
            t := strconv.Itoa(k)
            wt.WriteString(t + "\n")
        }
    }
}

func d(n int) (r int) {
    r = n
    for n != 0 {
        r += n % 10
        n /= 10
    }

    return
}
