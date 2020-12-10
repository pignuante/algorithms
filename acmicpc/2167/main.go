package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    flag := 1
    for _, c := range sc.Bytes() {
        if c == '-' {
            flag = -1
            continue
        }
        r *= 10
        r += int(c - '0')
    }
    r *= flag
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    Y, X := nextInt(), nextInt()
    arr := make([][]int, Y, Y)
    for y := 0; y < Y; y++ {
        x_ := make([]int, X, X)
        for x := 0; x < X; x++ {
            x_[x] = nextInt()
        }
        arr[y] = x_
    }
    K := nextInt()

    for ; K > 0; K-- {
        i, j, x, y := nextInt(), nextInt(), nextInt(), nextInt()
        sum := 0
        for xx := i - 1; xx < x; xx++ {
            for yy := j - 1; yy < y; yy++ {
                sum += arr[xx][yy]
            }
        }
        wt.WriteString(strconv.Itoa(sum))
        wt.WriteByte('\n')
    }
}
