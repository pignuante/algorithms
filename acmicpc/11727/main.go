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

	n := nextInt()
	wt.WriteString(strconv.Itoa(Tiling(n)))
}

func Tiling(n int) int {
	a, b := 1, 3
	for i := 1; i < n; i++ {
		a, b = b, (b+(a*2))%10007
	}
	return a
}
