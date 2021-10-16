package main

import (
	"bufio"
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
	isRun := true
	for isRun {
		sc.Scan()
		t := sc.Text()
		if t[0] == '0' {
			break
		}
		if IsPalindrome(t) {
			wt.WriteString("yes")
		} else {
			wt.WriteString("no")
		}
		wt.WriteByte('\n')
	}

}

func IsPalindrome(str string) bool {
	l := len(str)
	if l < 2 {
		return true
	}
	if str[0] != str[l-1] {
		return false
	}

	return IsPalindrome(str[1 : l-1])
}
