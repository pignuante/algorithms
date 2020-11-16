package main

import (
    "os"
    "bufio"
    "strconv"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 5)
    wr := bufio.NewWriterSize(os.Stdout, 5)
    defer wr.Flush()

    line, _, _ := rd.ReadLine()
    value := 0
    for _, v := range line {
        value *= 10
        value += int(v-'0')
    }
    for i := 1; i <= value; i++ {
        wr.WriteString(strconv.Itoa(i))
        wr.WriteByte('\n')
    }

}
