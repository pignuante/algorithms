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
	for _, c := range sc.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	N, M := nextInt(), nextInt()
	notListen := make(map[string]bool)
	var result []string

	for ; N > 0; N-- {
		sc.Scan()
		notListen[sc.Text()] = true
	}
	for ; M > 0; M-- {
		sc.Scan()
		if name := sc.Text(); notListen[name] {
			result = append(result, name)
		}
	}
	sort.Strings(result)
	wt.WriteString(strconv.Itoa(len(result)) + "\n")
	for _, value := range result {
		wt.WriteString(value)
		wt.WriteByte('\n')
	}
}
