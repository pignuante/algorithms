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
    message := ""
    stack := make([]int, 100000)
    for count := 1; N > 0; N-- {
        sc.Scan()
        n, _ := strconv.Atoi(sc.Text())

        for n >= count {
            stack = append(stack, count)
            message += "+\n"
            count++
        }
        if stack[len(stack)-1] != n {
            wt.WriteString("NO")
            return
        }
        message += "-\n"
        stack = stack[:len(stack)-1]
    }
    wt.WriteString(message)
}
