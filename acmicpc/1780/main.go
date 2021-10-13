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

type Data struct {
	n, z, p, a int
}

func main() {
	sc.Split(bufio.ScanWords)
	size := nextInt()
	paper := make([][]int8, size)
	for i := range paper {
		paper[i] = make([]int8, size)
		for j := range paper[i] {
			paper[i][j] = int8(nextInt())
		}
	}

	result := Search(0, 0, size, &paper)

	fmt.Printf("%d\n%d\n%d\n", result.n, result.z, result.p)
}
func Search(y, x, size int, paper *[][]int8) Data {
	if size == 1 {
		switch (*paper)[y][x] {
		case -1:
			return Data{1, 0, 0, -1}
		case 0:
			return Data{0, 1, 0, 0}
		case 1:
			return Data{0, 0, 1, 1}
		default:
			panic(0)
		}
	}
	t := size / 3
	m1, m2, m3 := Search(y, x, t, paper), Search(y, x+t, t, paper), Search(y, x+t*2, t, paper)
	m4, m5, m6 := Search(y+t, x, t, paper), Search(y+t, x+t, t, paper), Search(y+t, x+t*2, t, paper)
	m7, m8, m9 := Search(y+t*2, x, t, paper), Search(y+t*2, x+t, t, paper), Search(y+t*2, x+t*2, t, paper)
	if m1.a != 2 && m1.a == m2.a && m1.a == m3.a && m1.a == m4.a && m1.a == m5.a && m1.a == m6.a && m1.a == m7.a && m1.a == m8.a && m1.a == m9.a {
		return m1
	}
	return Data{
		m1.n + m2.n + m3.n + m4.n + m5.n + m6.n + m7.n + m8.n + m9.n,
		m1.z + m2.z + m3.z + m4.z + m5.z + m6.z + m7.z + m8.z + m9.z,
		m1.p + m2.p + m3.p + m4.p + m5.p + m6.p + m7.p + m8.p + m9.p,
		2,
	}
}
