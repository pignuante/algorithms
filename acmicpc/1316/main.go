package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    count := 0
    for i := 0; i < N; i++ {
        sc.Scan()
        word := sc.Text()
        if isCheckedGroupWord(word) {
            count++
        }
    }
    wt.WriteString(strconv.Itoa(count))
}

func isCheckedGroupWord(word string) bool {
    checker := make([]bool, 26)
    flag := false
    var prev byte

    for _, v := range word {
        idx := byte(v - 'a')
        if checker[idx] && prev != idx {
            flag = false
            return flag
        }
        checker[idx] = true
        prev = idx
    }
    return true
}
