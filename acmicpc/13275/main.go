package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()
	sc.Scan()
	temp := sc.Text()
	str := ""
	for _, v := range temp {
		str += "#"
		str += string(v)
	}
	str += "#"
	Manahcers(str, len(str))

	l := len(str)
	ans := 1
	for i := 0; i < l; i++ {
		ans = Max(ans, A[i])
	}
	wt.WriteString(strconv.Itoa(ans))
}

var temp, str string

const MAXN int = 100001 * 2

var A []int = make([]int, MAXN, MAXN)

func Manahcers(str string, N int) {
	r, p := 0, 0
	for i := 0; i < N; i++ {
		if i <= r {
			A[i] = Min(A[2*p-i], r-i)
		} else {
			A[i] = 0
		}
		for i-A[i]-1 >= 0 && i+A[i]+1 < N && str[i-A[i]-1] == str[i+A[i]+1] {
			A[i]++
		}
		if r < i+A[i] {
			r = i + A[i]
			p = i
		}
	}
}
func Max(a, b int) (max int) {
	max = a
	if a < b {
		max = b
	}
	return
}
func Min(a, b int) (min int) {
	min = a
	if a > b {
		min = b
	}
	return
}
