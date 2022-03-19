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
	number := nextInt()
	result := Factorization(number)

	wt.WriteString(strconv.Itoa(result))
}

func Factorization(number int) int {
	result := 0
	for i := 1; i <= number; i++ {
		temp := i
		for temp%5 == 0 && temp > 0 {
			result++
			temp /= 5
		}
	}
	return result
}
