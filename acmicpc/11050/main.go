package main

import (
	"bufio"
	"os"
	"strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	N, K := nextInt(), nextInt()
	wt.WriteString(strconv.Itoa(Binomial(N, K)))
}

func Binomial(N, K int) int {
	return Factorial(N) / (Factorial(K) * Factorial(N-K))
}
func Factorial(num int) int {
	if num == 1 || num == 0 {
		return 1
	}
	return Factorial(num-1) * num
}
