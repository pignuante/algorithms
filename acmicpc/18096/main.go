package main

import (
    "bufio"
    "os"
    // "strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer wt.Flush()
    wt.WriteString("1")
}
