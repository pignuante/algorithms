package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    a, b := 0, 0
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)

    for true {
        fmt.Fscanln(reader, &a, &b)
        if a == 0 && b == 0 {
            break
        }
        fmt.Fprintln(writer, a+b)
    }
    writer.Flush()
}
