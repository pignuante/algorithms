package main

import (
    "bufio"
    "os"
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

    for {
        auset, ausar, heru := nextInt(), nextInt(), nextInt()
        if auset == 0 && ausar == 0 && heru == 0 {
            break
        }
        if heru < auset {
            heru, auset = auset, heru
        }
        if heru < ausar {
            heru, ausar = ausar, heru
        }
        if auset*auset+ausar*ausar == heru*heru {
            wt.WriteString("right\n")
        } else {
            wt.WriteString("wrong\n")
        }
    }
}
