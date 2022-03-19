package main

import (
	"bufio"
	"os"
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
	N, M := nextInt(), nextInt()
	passwords := make(map[string]string)
	for ; N > 0; N-- {
		sc.Scan()
		address := sc.Text()
		sc.Scan()
		password := sc.Text()
		passwords[address] = password
	}
	for ; M > 0; M-- {
		sc.Scan()
		address := sc.Text()
		wt.WriteString(passwords[address])
		wt.WriteByte('\n')
	}
}
