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

    girl, boy, intern := nextInt(), nextInt(), nextInt()
    result := 0
    for people := girl + boy - intern; girl > 1 && boy > 0 && people > 2; {
        girl -= 2
        boy -= 1
        people -= 3
        result++
    }
    wt.WriteString(strconv.Itoa(result))
}
