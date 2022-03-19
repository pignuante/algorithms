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
	N := nextInt()
	for ; N > 0; N-- {
		n := nextInt()
		clothes := make(map[string]int)
		for ; n > 0; n-- {
			sc.Scan()
			_ = sc.Text()
			sc.Scan()
			kind := sc.Text()
			clothes[kind]++
		}
		nums := 1

		for _, value := range clothes {
			nums *= value + 1
		}
		wt.WriteString(strconv.Itoa(nums - 1))
		wt.WriteByte('\n')
	}
}
