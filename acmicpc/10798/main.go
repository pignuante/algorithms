package main

import (
    "bufio"
    "os"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    lines, max := 5, 0
    strs := make([][15]byte, lines, lines)
    for line := 0; line < lines; line++ {
        sc.Scan()
        t := sc.Bytes()
        if max < len(t) {
            max = len(t)
        }
        for k, char := range t {
            strs[line][k] = char
        }
    }
    for x := 0; x < max; x++ {
        for y := 0; y < lines; y++ {
            if  strs[y][x] == 0 {
                continue
            }
            wt.WriteByte(strs[y][x])
        }
    }
}
