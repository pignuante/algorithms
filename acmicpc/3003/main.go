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

var Chess = map[int]int{
	0: 1, 1: 1, 2: 2, 3: 2, 4: 2, 5: 8,
}

func main() {
	sc.Split(bufio.ScanWords)
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()
	for i := 0; i < 6; i++ {
		wt.WriteString(strconv.Itoa(Chess[i] - nextInt()))
		wt.WriteByte(' ')
	}

}
