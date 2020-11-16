package main

import (
    "os"
    "bufio"
    "strconv"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 5)

    line, _, _ := rd.ReadLine()
    value, _ := strconv.Atoi(string(line))
    wr := bufio.NewWriterSize(os.Stdout, 5 * value)
    defer wr.Flush()

    for i := value; i > 0 ; i-- {
        wr.WriteString(strconv.Itoa(i))
        wr.WriteByte('\n')
    }

}
