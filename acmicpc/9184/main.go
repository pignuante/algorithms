package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (res int64) {
	sc.Scan()
	text := sc.Text()
	//v, _ := strconv.Atoi(text)
	v, _ := strconv.ParseInt(text, 10, 64)
	return v
}

func main() {
	const size = 21
	sc.Split(bufio.ScanWords)
	defer wt.Flush()

	table := make([][][]int64, size)
	for z := range table {
		table[z] = make([][]int64, size)
		for y := range table[z] {
			table[z][y] = make([]int64, size)
			for x := range table[z][y] {
				table[z][y][x] = -1
			}
		}
	}

	for {
		a, b, c := nextInt(), nextInt(), nextInt()
		if a == -1 && b == -1 && c == -1 {
			break
		}
		result := w(&table, a, b, c)
		fmt.Fprintf(wt, "w(%d, %d, %d) = %d\n", a, b, c, result)
	}
}

func w(table *[][][]int64, a, b, c int64) int64 {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}
	if a > 20 || b > 20 || c > 20 {
		return w(table, 20, 20, 20)
	}
	if temp := (*table)[a][b][c]; temp != -1 {
		return temp
	}
	if a < b && b < c {
		(*table)[a][b][c] = w(table, a, b, c-1) + w(table, a, b-1, c-1) - w(table, a, b-1, c)
	}
	(*table)[a][b][c] = w(table, a-1, b, c) + w(table, a-1, b-1, c) + w(table, a-1, b, c-1) - w(table, a-1, b-1, c-1)
	return (*table)[a][b][c]
}
