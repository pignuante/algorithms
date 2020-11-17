package main

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    rd := bufio.NewReader(os.Stdin)
    wt := bufio.NewWriter(os.Stdout)

    defer wt.Flush()

    line ,_, _ := rd.ReadLine()
    t := strings.Split(string(line), " ")
    x, _ := strconv.Atoi(t[1])

    line2, _, _ := rd.ReadLine()
    values := strings.Split(string(line2), " ")
    for _, v := range values {
        t, _ := strconv.Atoi(string(v))
        if t < x {
            wt.WriteString(string(v) + " ")
        }
    }

}
