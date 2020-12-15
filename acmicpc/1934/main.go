package main

import (
	"bufio"
	"os"
	"strconv"
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

	T := nextInt()
	for ; T > 0; T-- {
		a, b := nextInt(), nextInt()
		lcm := LCM(a, b, GCD(a, b))
		wt.WriteString(strconv.Itoa(lcm))
		wt.WriteByte('\n')
	}
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b, r int) int {
	return a * b / r
}
