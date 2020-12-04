package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r = 0
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

    T := nextInt()
    for ; T > 0; T-- {
        x, y := nextInt(), nextInt()
        wt.WriteString(strconv.Itoa(getMoveCount(y-x)) + "\n")
    }

}

func getMoveCount(dist int) (count int) {
    root := findRoot(dist)
    if dist == root*root {
        count = 2*root - 1
    } else if dist <= root*root+root {
        count = 2 * root
    } else {
        count = 2*root + 1
    }
    return
}

func findRoot(dist int) (root int) {
    root = 1
    for {
        if dist >= root*root && dist < (root+1)*(root+1) {
            break
        }
        root++
    }
    return
}
