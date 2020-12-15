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
    sticks := sc.Bytes()
    var stack []byte
    sum := 0
    for k, v := range sticks {
        if v == '(' {
            stack = append(stack, v)
        } else {
            stack = stack[1:]
            if sticks[k-1] == '(' {
                sum += len(stack)
            } else {
                sum++
            }
        }
    }
    wt.WriteString(strconv.Itoa(sum))
}
