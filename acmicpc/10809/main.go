package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

const num int = 26

func main() {
    var alpabets [num]int

    defer wt.Flush()

    sc.Scan()
    S := sc.Text()

    initArray(&alpabets)
    findIndex(S, &alpabets)
    printResult(alpabets)

}

func printResult(alpabets [num]int) {
    for _, v := range alpabets {
        wt.WriteString(strconv.Itoa(v) + " ")
    }
}

func findIndex(S string, alpabets *[num]int) {
    for k, v := range S {
        idx := int(v) - int('a')
        if alpabets[idx] == -1 {
            alpabets[idx] = k
        }
    }
}

func initArray(arr *[num]int) {
    for k, _ := range arr {
        arr[k] = -1
    }
}
