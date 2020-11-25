package main

import (
    "bufio"
    "os"
    "strconv"
)
func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()
    sc.Scan()
    T, _ := strconv.Atoi(sc.Text())
    for ; T > 0; T-- {
        sc.Scan()
        PS := sc.Text()
        counter := 0
        for _, v := range PS {
            if v == '(' {
                counter++
            } else {
                counter--
            }
            if counter < 0 {
                break
            }
        }
        if counter == 0 {
            wt.WriteString("YES\n")
        } else {
            wt.WriteString("NO\n")
        }
    }
}
