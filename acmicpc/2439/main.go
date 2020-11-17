package main

import (
    "os"
    "bufio"
    "strconv"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 5)

    line, _, _ := rd.ReadLine()
    n, _ :=  strconv.Atoi(string(line))

    rw := bufio.NewWriterSize(os.Stdout, 5*n)
    defer rw.Flush()

    spaces := make([]byte, n)

    for i := 0; i < n; i++ {
        spaces[i] = ' '
    }

    for i := n-1; i >= 0; i-- {
        spaces[i] = '*'
        rw.WriteString(string(spaces))
        rw.WriteByte('\n')
    }

}
