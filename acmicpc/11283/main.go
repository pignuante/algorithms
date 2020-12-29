package main

import (
	"bufio"
	"os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()

	sc.Scan()
	t := sc.Text()

	for _, v := range t {
		wt.WriteString(strconv.Itoa(int(v) - 44031))
	}
}
