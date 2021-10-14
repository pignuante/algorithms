package main

import (
	"bufio"
	"fmt"
	"os"
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
	N, r, c := nextInt(), nextInt(), nextInt()
	count, result := 0, 0
	Solve(N, c, r, 0, 0, Pow(2, N), &count, &result)
	fmt.Println(result)

}

func Solve(N, c, r, y, x, size int, cnt, result *int) {
	if r == y && c == x {
		*result = *cnt
		return
	}
	if r < y && r >= y && c < x+N && c >= x {
		t := size / 2
		Solve(N, c, r, y, x, t, cnt, result)
		Solve(N, c, r, y, x+t, t, cnt, result)
		Solve(N, c, r, y+t, x, t, cnt, result)
		Solve(N, c, r, y+t, x+t, t, cnt, result)
	} else {
		*cnt += N * N
	}
}

func Pow(number, times int) (result int) {
	result = number
	for i := 0; i < times; i++ {
		result *= number
	}
	return
}
