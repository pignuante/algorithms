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

	N := nextInt()
	rooms := make([]int64, N, N)
	for i := 0; i < N; i++ {
		rooms[i] = int64(nextInt())
	}

	B, C, viewer := int64(nextInt()), int64(nextInt()), int64(0)
	for k := range rooms {
		rooms[k] = Max(rooms[k]-B, 0)
		viewer++
		if rooms[k]%C == 0 {
			viewer += (rooms[k] / C)
		} else {
			viewer += (rooms[k] / C) + 1
		}
	}
	wt.WriteString(strconv.FormatInt(viewer, 10))
}

func Max(a, b int64) (max int64) {
	max = a
	if a < b {
		max = b
	}
	return
}
