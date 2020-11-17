package main

import (
    "os"
    "bufio"
    "strconv"
    // "fmt"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 5)

    line, _, _ := rd.ReadLine()
    num, _ := strconv.Atoi(string(line))
    wr := bufio.NewWriterSize(os.Stdout, 5 * num)
    defer wr.Flush()

    for i := 1; i <= num; i++ {
        line, _, _ = rd.ReadLine()
        a, b := 0, 0
        for _, v := range line {
            if  v == ' ' {
                a = b
                b = 0
                continue
            }
            b *= 10
            b += int(v-'0')
        }
        wr.WriteString("Case #")

        wr.WriteString(strconv.Itoa(i))
        wr.WriteString(": ")
        wr.WriteString(strconv.Itoa(a+b))
        wr.WriteByte('\n')
    }

}
