package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    text := sc.Text()
    text = repalceCroatian(text)
    l := strconv.Itoa(len(text))
    wt.WriteString(l)

}

func repalceCroatian(text string) string {
    croatians := []string{"dz=", "c=", "c-", "d-", "lj", "nj", "s=", "z="}
    for _, v := range croatians {
        text = strings.Replace(text, v, "*", -1)
    }
    return text
}
