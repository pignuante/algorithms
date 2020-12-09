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
    alphabet := make([]int, 26, 26)
    for _, v := range sc.Bytes() {
        alphabet[v-'a']++
    }
    for _ , v := range alphabet {
        wt.WriteString(strconv.Itoa(v))
        wt.WriteByte(' ')
    }
}
