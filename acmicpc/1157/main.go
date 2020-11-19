package main

import (
    "bufio"
    "os"
    "strings"
)

var num int = 26

func main() {
    rd := bufio.NewReader(os.Stdin)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    var arr [26]int
    initArray(&arr)

    words, _ := rd.ReadString('\n')
    words = strings.Trim(strings.ToUpper(words), "\n")

    wordCounter(words, &arr)
    m := maxCounter(arr)
    wt.WriteString(m)
}

func printResult(m int) {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    if m == -1 {
        wt.WriteString("?")
    } else {
        wt.WriteString(string(m + int('A')))
    }
}

func maxCounter(arr [26]int) string {
    max := -1
    key := "?"
    for k, v := range arr {
        if v > max {
            key = string(k + int('a'))
            max = v
        } else if v == max {
            key = "?"
        }
    }
    return strings.ToUpper(key)
}

func wordCounter(words string, arr *[26]int) {
    for _, v := range words {
        idx := int(v) - int('A')
        arr[idx] += 1
    }
}

func initArray(arr *[26]int) {
    for k, _ := range arr {
        arr[k] = 0
    }
}
