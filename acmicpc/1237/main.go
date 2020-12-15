package main

import (
	"bufio"
	"os"
	// "strconv"
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
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()
	wt.WriteString("문제의 정답")
}
