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
	N, M := nextInt(), nextInt()
	nums := make([]int, N+1, N+1)

	nums[1] = nextInt()
	for i := 2; i <= N; i++ {
		nums[i] += nums[i-1] + nextInt()
	}
	for i := 0; i < M; i++ {
		start, end := nextInt(), nextInt()
		result := nums[end] - nums[start-1]
		wt.WriteString(strconv.Itoa(result))
		wt.WriteByte('\n')
	}
}
