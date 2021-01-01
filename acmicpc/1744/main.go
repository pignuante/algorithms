package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
	sc.Scan()
	flag := 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			flag = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	r *= flag
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()

	N, result := nextInt(), 0
	negative, positive := []int{}, []int{}
	for i := 0; i < N; i++ {
		if t := nextInt(); t <= 0 {
			negative = append(negative, t)
		} else if t == 1 {
			result += 1
		} else {
			positive = append(positive, t)
		}
	}
	sort.Ints(negative)
	sort.Slice(positive, func(a, b int) bool {
		return positive[a] > positive[b]
	})
	for i := 0; i < len(negative); i += 2 {
		if i+1 < len(negative) {
			result += (negative[i] * negative[i+1])
		} else {
			result += negative[i]
		}
	}
	for i := 0; i < len(positive); i += 2 {
		if i+1 < len(positive) {
			result += (positive[i] * positive[i+1])
		} else {
			result += positive[i]
		}
	}
	wt.WriteString(strconv.Itoa(result))
}
