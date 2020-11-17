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

    for {
        strs, err := rd.ReadString('\n')
        if err != nil {
           break
        }

        strTkns := strings.Split(strings.TrimSpace(string(strs)), " ")
        a, _ := strconv.Atoi(strTkns[0])
        b, _ := strconv.Atoi(strTkns[1])
        wt.WriteString(strconv.Itoa(a+b) + "\n")
    }

}
