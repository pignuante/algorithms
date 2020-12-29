package main

import (
	"bufio"
	"os"
	"strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
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
	defer wt.Flush()

	const tax = 0.22
	price := float64(nextInt())

	wt.WriteString(strconv.Itoa(int(price * (1 - tax))))
	wt.WriteByte(' ')
	wt.WriteString(strconv.Itoa(int(price*0.8 + price*0.2*(1-tax))))
}
