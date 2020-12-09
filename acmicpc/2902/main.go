package main

import (
    "bufio"
    "os"
    "strings"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    names := strings.Split(sc.Text(), "-")
    shorten := ""
    for _, v := range names {
        shorten += string(v[0])
    }
    wt.WriteString(shorten)
}
