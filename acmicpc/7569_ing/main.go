package main

import (
	"bufio"
	"fmt"
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

type Mvs struct {
	Z, Y, X int
}

var mvs = []Mvs{
	{0, 1, 0}, {0, -1, 0}, // 상하
	{0, 0, -1}, {0, 0, 1}, // 좌우
	{-1, 0, 0}, {1, 0, 0}, // 앞뒤
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	X, Y, Z := nextInt(), nextInt(), nextInt()
	box := make([][][]int, Z, Z)
	for z := 0; z < Z; z++ {
		box[z] = make([][]int, Y, Y)
		for y := 0; y < Y; y++ {
			box[z][y] = make([]int, X, X)
			for x := 0; x < X; x++ {
				box[z][y][x] = nextInt()
			}
		}

	}

	fmt.Println(box)
}
