package main

import (
	"bufio"
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
	commands := nextInt()
	bits := make([]bool, 21, 21)
	for ; commands > 0; commands-- {
		sc.Scan()
		command := sc.Text()
		Check(command, &bits)
	}

}

func Check(command string, bits *[]bool) {
	switch command {
	case "add":
		x := nextInt()
		(*bits)[x] = true
	case "remove":
		x := nextInt()
		(*bits)[x] = false
	case "check":
		x := nextInt()
		if (*bits)[x] {
			wt.WriteByte('1')
		} else {
			wt.WriteByte('0')
		}
		wt.WriteByte('\n')
	case "toggle":
		x := nextInt()
		(*bits)[x] = !(*bits)[x]
	case "all":
		for idx := range *bits {
			(*bits)[idx] = true
		}
	case "empty":
		for idx := range *bits {
			(*bits)[idx] = false
		}
	}
}
