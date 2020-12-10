package main

import (
    "bufio"
    "os"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    wt.WriteString(`       _.-;;-._
'-..-'|   ||   |
'-..-'|_.-;;-._|
'-..-'|   ||   |
'-..-'|_.-''-._|`)
}
