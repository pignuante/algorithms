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

	num := nextInt()
	Factorization(num)
}

func Factorization(num int) {
	idx := 2
	for num != 1 {
		if num%idx == 0 {
			num /= idx
			wt.WriteString(strconv.Itoa(idx))
			wt.WriteByte('\n')
		} else {
			idx += 1
		}
	}
}
