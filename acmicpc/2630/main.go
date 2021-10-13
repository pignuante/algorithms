package main

import (
	"bufio"
	"fmt"
	"os"
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
	N := nextInt()
	confetti := make([][]bool, N, N)
	for y := 0; y < N; y++ {
		confetti[y] = make([]bool, N, N)
		for x := 0; x < N; x++ {
			if t := nextInt(); t == 0 {
				confetti[y][x] = false
			} else {
				confetti[y][x] = true
			}
		}
	}
	a, b := DFS(0, 0, N, &confetti)
	fmt.Println(a, b)

}

func DFS(y, x, length int, confetti *[][]bool) (white, blue int) {
	if length == 1 {
		if (*confetti)[y][x] {
			white, blue = 0, 1
			return
		} else {
			white, blue = 1, 0
			return
		}
	}

	temp := length / 2
	w0, b0 := DFS(y, x, temp, confetti)
	w1, b1 := DFS(y, x+temp, temp, confetti)
	w2, b2 := DFS(y+temp, x, temp, confetti)
	w3, b3 := DFS(y+temp, x+temp, temp, confetti)
	w := w0 + w1 + w2 + w3
	b := b0 + b1 + b2 + b3
	if w == 0 {
		white, blue = 0, 1
		return
	}
	if b == 0 {
		white, blue = 1, 0
		return
	}
	return w, b
}
