package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {

    sc := bufio.NewScanner(os.Stdin)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    for i := 0; i < N; i++ {
        var OX string
        if sc.Scan() {
            OX = sc.Text()
        }
        score := 0
        temp := 0
        for _, v := range OX {
            if string(v) == "X" {
                temp = 0
            } else {
                temp++
                score += temp
            }
        }
        wt.WriteString(strconv.Itoa(score) + "\n")
    }
}
