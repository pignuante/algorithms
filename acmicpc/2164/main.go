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
    N, _ := strconv.Atoi(sc.Text())
    queue := make([]int, N, N+1)
    for k := range queue {
        queue[k] = k+1
    }

    for len(queue) > 1 {
        queue = append(queue[2:], queue[1])
    }
    wt.WriteString(strconv.Itoa(queue[0]))
}
