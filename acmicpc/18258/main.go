package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    var queue []int
    for ; N > 0; N-- {
        sc.Scan()
        command := strings.Fields(sc.Text())
        switch command[0] {
        case "push":
            num, _ := strconv.Atoi(command[1])
            queue = append(queue, num)
        case "pop":
            num := -1
            if len(queue) > 0 {
                num = queue[0]
                queue = queue[1:]
            }
            wt.WriteString(strconv.Itoa(num) + "\n")
        case "size":
            wt.WriteString(strconv.Itoa(len(queue)) + "\n")
        case "empty":
            num := 0
            if len(queue) == 0 {
                num = 1
            }
            wt.WriteString(strconv.Itoa(num) + "\n")
        case "front":
            num := -1
            if len(queue) > 0 {
                num = queue[0]
            }
            wt.WriteString(strconv.Itoa(num) + "\n")
        case "back":
            num := -1
            if len(queue) > 0 {
                num = queue[len(queue)-1]
            }
            wt.WriteString(strconv.Itoa(num) + "\n")
        }
    }
}
