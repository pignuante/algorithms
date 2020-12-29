package main

import (
	"bufio"
	"math/big"
	"os"
	// "strconv"
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

	var a, b, c big.Int
	sc.Scan()
	a.SetString(sc.Text(), 10)
	sc.Scan()
	b.SetString(sc.Text(), 10)
	sc.Scan()
	c.SetString(sc.Text(), 10)
	t := a.Add(&a, &b)
	wt.WriteString(c.Add(&c, t).String())
}
