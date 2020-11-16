package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    rd := bufio.NewReaderSize(os.Stdin, 1<<24)
    wr := bufio.NewWriterSize(os.Stdout, 1<<23)
    defer wr.Flush()

    rd.ReadLine()

    for {
        line, _, err := rd.ReadLine()
        if err != nil {
            break
        }
        ret := calc(line)
        wr.WriteString(strconv.Itoa(ret))
        wr.WriteByte('\n')
    }
}

func calc(line []byte) int {
    tmp := 0
    ans := 0
    for _, v := range line {
        if v == ' ' {
            ans = tmp
            tmp = 0
            continue
        }

        tmp *= 10
        tmp += int(v-'0')
    }

    return tmp+ans
}
