package main

import (
    "bufio"
    "math/big"
    "os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

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
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    var n, m big.Int
    sc.Scan()
    n.SetString(sc.Text(), 10)
    sc.Scan()
    m.SetString(sc.Text(), 10)

    q := big.NewInt(0).Div(&n, &m)
    r := big.NewInt(0).Mod(&n, &m)
    wt.WriteString(q.String())
    wt.WriteByte('\n')
    wt.WriteString(r.String())
}
