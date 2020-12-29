package main

import (
    "bufio"
    "os"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    sc.Scan()
    _ = sc.Text()
    wt.WriteString("1")
}
