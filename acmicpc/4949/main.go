package main

import (
    "bufio"
    "os"
    // "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanLines)

    defer wt.Flush()
    end := "."
    for {
        sc.Scan()
        t := sc.Text()
        if t == end {
            break
        }
        parentheses := 0
        squareBracket := 0
        for _, v := range t{
            if v == '(' {
                parentheses++
            } else if v =='[' {
                squareBracket++
            } else if v == ')' {
                parentheses--
            } else if v ==']' {
                squareBracket--
            }
            if parentheses < 0 {
                break
            }
        }
        if parentheses == 0 && squareBracket == 0 {
            wt.WriteString("yes\n")
        } else {
            wt.WriteString("no\n")
        }
    }
}
