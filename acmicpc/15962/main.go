package main

import (
	"bufio"
	"os"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer wt.Flush()
	wt.WriteString("파이팅!!")
}
