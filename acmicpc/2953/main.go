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

    idx, sum := -1, 0
    for y := 0; y < 5; y++ {
        s := nextInt() + nextInt() + nextInt() + nextInt()
        if s > sum {
            idx, sum = y, s
        }
    }
    wt.WriteString(strconv.Itoa(idx+1) + " " + strconv.Itoa(sum))

}
