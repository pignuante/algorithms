package main

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 100000)

    line ,_, _ := rd.ReadLine()
    t := strings.Split(string(line), " ")
    x, _ := strconv.Atoi(t[1])

    wt := bufio.NewWriterSize(os.Stdout, 100000)
    defer wt.Flush()

    line2, _, _ := rd.ReadLine()
    values := strings.Split(string(line2), " ")

    for _, v := range values {
        t, _ := strconv.Atoi(v)
        if t < x {
            wt.WriteString(v + " ")
        }
    }

}
