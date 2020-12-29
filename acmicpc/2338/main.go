package main

import (
	"bufio"
	"math/big"
	"os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()

	var A, B big.Int
	sc.Scan()
	A.SetString(sc.Text(), 10)
	sc.Scan()
	B.SetString(sc.Text(), 10)

	plus := big.NewInt(0).Add(&A, &B).String()
	minus := big.NewInt(0).Sub(&A, &B).String()
	mul := big.NewInt(0).Mul(&A, &B).String()

	wt.WriteString(plus)
	wt.WriteByte('\n')
	wt.WriteString(minus)
	wt.WriteByte('\n')
	wt.WriteString(mul)
	wt.WriteByte('\n')
}