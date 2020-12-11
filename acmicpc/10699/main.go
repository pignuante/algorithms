package main

import (
	"bufio"
	"os"
	"time"
)

func main() {
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()
	t := time.Now().UTC()
	t = t.In(time.FixedZone("Asia/Seoul", 9*60*60))
	wt.WriteString(t.Format("2006-01-02"))
}
