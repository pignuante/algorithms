package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    sc.Scan()
    line := sc.Text()
    plus := strings.Split(line, "-")
    fmt.Println(minus(sum(plus)))
}

func minus(minus []int) (r int) {
    r = minus[0]
    for _, v := range minus[1:] {
        r -= v
    }
    return
}
func sum(plus []string) []int {
    r := make([]int, len(plus), len(plus))
    for k, v := range plus {
        t := strings.Split(v, "+")
        s := 0
        for _, vv := range t {
            tt, _ := strconv.Atoi(vv)
            s += tt
        }
        r[k] += s
    }
    return r
}
