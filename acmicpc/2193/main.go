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
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	wt.WriteString(strconv.Itoa(a))

}
