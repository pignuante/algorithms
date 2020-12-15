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

	cards := nextInt()
	deck := make([]int, cards+1, cards+1)
	for card := 1; card <= cards; card++ {
		deck[card] = nextInt()
	}
	wt.WriteString(strconv.Itoa(Calc(&deck, cards)))
}

func Calc(deck *[]int, cards int) int {
	dp := make([]int, cards+1, cards+1)
	for i := 1; i <= cards; i++ {
		for k := 1; k <= i; k++ {
			dp[i] = Max(dp[i], dp[i-k]+(*deck)[k])
		}
	}
	return dp[cards]
}

func Max(a, b int) (max int) {
	max = a
	if a < b {
		max = b
	}
	return
}
