package main

import (
    "os"
    "bufio"
    "strconv"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 5)

    line, _, _ := rd.ReadLine()
    n, _ := strconv.Atoi(string(line))

    rw := bufio.NewWriterSize(os.Stdout, 5*n)
    defer rw.Flush()

    var star string = "*"
    for i := 0; i < n; i++ {
        rw.WriteString(star)
        rw.WriteByte('\n')
        star += "*"
    }

}
