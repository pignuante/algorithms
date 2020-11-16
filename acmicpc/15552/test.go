package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    br := bufio.NewReader(os.Stdin)
    bw := bufio.NewWriter(os.Stdout)
    defer bw.Flush()

    line, _, _ := br.ReadLine()

    tmp := 0
    ans := 0
    fmt.Println(line)
    for _, v := range line {
        if v == ' ' {
            ans = tmp
            tmp = 0
            continue
        }
        tmp *= 10
        fmt.Println(tmp)
        tmp += int(v-'0')

        fmt.Println(1, int(v), tmp)
    }

    fmt.Println(tmp, "asdf", ans)

}

