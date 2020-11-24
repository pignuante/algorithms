package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	T, _ := strconv.Atoi(sc.Text())
	for ; T > 0; T-- {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		wt.WriteString(strconv.Itoa(P(n)) + "\n")
	}
}

func P(n int) int {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 5, 7, 9}
	a, b, c := 5, 7, 9
	if n < 10 {
		return nums[n-1]
	}
	for ; n > 10; n-- {
		a, b, c = b, c, a+b
	}

	return c
}
