package main

import (
    "bufio"
    "fmt"
    "math/big"
    "os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
    defer w.Flush()
    a, b := new(big.Int), new(big.Int)

    fmt.Fscanf(r, "%d %d\n", a, b)
    fmt.Fprintln(w, a.Mul(a, b))
}
