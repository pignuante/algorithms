package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    wt := bufio.NewWriter(os.Stdout)
    sc.Split(bufio.ScanLines)
    defer wt.Flush()

    var T int = 0
    if sc.Scan() {
        T, _ = strconv.Atoi(sc.Text())
    }

    for i := 0; i < T; i++ {
        sc.Scan()
        str := sc.Text()
        rs := strings.Split(str, " ")
        R, _ := strconv.Atoi(rs[0])
        S := rs[1]
        for _, v := range S {
            for i := 0; i < R; i++ {
                wt.WriteString(string(v))
            }
        }
        wt.WriteByte('\n')
    }
}
