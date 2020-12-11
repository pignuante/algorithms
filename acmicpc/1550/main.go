package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var wt  = bufio.NewWriter(os.Stdout)
func main() {
    defer wt.Flush()
    sc.Scan()
    hexadecimal := sc.Text()
    wt.WriteString(strconv.Itoa(HexToDec(hexadecimal)))
}

func HexToDec(hex string) (r int){
    for k, v := range hex {
        r *= 16
        if 'A' <= v && v <= 'F' {
            r += int(hex[k] - 55)
        } else {
            r += int(hex[k] - '0')
        }
    }
    return
}

