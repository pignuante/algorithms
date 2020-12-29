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

	var a, b big.Int
	sc.Scan()
	a.SetString(sc.Text(), 10)
	sc.Scan()
	b.SetString(sc.Text(), 10)

	r1 := big.NewInt(0).Add(&a, &b)
	r2 := big.NewInt(0).Sub(&a, &b)

	wt.WriteString(r1.Div(r1, big.NewInt(2)).String())
	wt.WriteByte('\n')
	wt.WriteString(r2.Div(r2, big.NewInt(2)).String())
}
