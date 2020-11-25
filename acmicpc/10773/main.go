package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()

    sc.Scan()
    K, _ := strconv.Atoi(sc.Text())
    var stack []int
    for ; K > 0; K-- {
        sc.Scan()
        n, _ := strconv.Atoi(sc.Text())
        if n == 0 {
            stack = stack[1:]
        } else {
            t := []int{n}
            stack = append(t, stack...)
        }
    }
    r := 0
    for _, v := range stack {
        r += v
    }
    wt.WriteString(strconv.Itoa(r))
}
