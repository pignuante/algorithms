package main

import (
    "bufio"
    "os"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    sc.Scan()
    N := sc.Text()
    for i := '9'; i >= '0'; i-- {
        for _, v := range N {
            if i == v {
                wt.WriteString(string(int(v)))
            }
        }
    }
}
