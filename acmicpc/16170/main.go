package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer wt.Flush()

	now := time.Now().UTC().In(time.FixedZone("KST", 9*60*60))
	
	fmt.Fprintf(wt, "%04d\n%02d\n%02d", now.Year(), now.Month(), now.Day())
}
