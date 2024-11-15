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
	const hongik = 1946
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	year := nextInt()
	wt.WriteString(strconv.Itoa(year - hongik))
}
