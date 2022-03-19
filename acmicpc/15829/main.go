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
	L := nextInt()
	sc.Scan()
	s := sc.Text()
	r, MOD, v := 31, 1234567891, 1
	result := 0
	for i := 0; i < L; i++ {
		result += (int(s[i]-'a'+1) * v) % MOD
		v = (v * r) % MOD
	}
	wt.WriteString(strconv.Itoa(result % MOD))
}
