package main

import (
    "bufio"
    "math"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    r, _ := strconv.ParseFloat(sc.Text(), 64)
    wt.WriteString(strconv.FormatFloat(r*r*math.Pi, 'f', 6, 64) + "\n" + strconv.FormatFloat(r*r*2, 'f', 6, 64))
}
