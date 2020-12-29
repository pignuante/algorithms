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

	Ur, Tr, Uo, To := nextInt(), nextInt(), nextInt(), nextInt()

	wt.WriteString(strconv.Itoa(56*Ur + 24*Tr + 14*Uo + 6*To))
}
