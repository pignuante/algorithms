package main

import (
	"bufio"
	"os"
	"sort"
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

	N := nextInt()
	A, B := make([]int, N, N), make([]int, N, N)

	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	for i := 0; i < N; i++ {
		B[i] = nextInt()
	}
	sort.Ints(A)
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})
	wt.WriteString(strconv.Itoa(Sum(&A, &B)))
}

func Sum(A, B *[]int) (sum int) {
	for k := range *A {
		sum += (*A)[k] * (*B)[k]
	}
	return
}
