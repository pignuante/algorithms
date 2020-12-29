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

type Cake struct {
	x, y, z int
}

func (a Cake) Sub(c Cake) (b Cake) {
	b.x = c.x - a.z
	b.y = c.y / a.y
	b.z = c.z - a.x
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()

	a := Cake{x: nextInt(), y: nextInt(), z: nextInt()}
	c := Cake{x: nextInt(), y: nextInt(), z: nextInt()}
	b := a.Sub(c)

	wt.WriteString(strconv.Itoa(b.x))
	wt.WriteByte(' ')
	wt.WriteString(strconv.Itoa(b.y))
	wt.WriteByte(' ')
	wt.WriteString(strconv.Itoa(b.z))
}
