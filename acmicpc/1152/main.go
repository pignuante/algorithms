package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 1000000)
    line, _, _ := rd.ReadLine()
    str := bytes.TrimSpace([]byte(line))

    if len(str) == 0 {
        fmt.Print(0)
        return
    }
    count := 1
    for _, char := range str {
        if char == byte(' ') {
            count++
        }
    }
    fmt.Print(count)
}
