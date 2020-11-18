package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanBytes)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    text := strconv.Itoa(int(sc.Bytes()[0]))

    wt.WriteString(text)

}
