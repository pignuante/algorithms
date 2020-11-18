package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanWords)

    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    d := sc.Text()
    N, _ := strconv.Atoi(d)

    var count int = 0
    for i := 1; i < N+1; i++ {
        if hansu(i) {
            count++
        }
    }
    wt.WriteString(strconv.Itoa(count))

}

func hansu(num int) bool {
    str := strconv.Itoa(num)
    if len(str) > 2 {
        minus := str[0] - str[1]
        for i := 1; i < len(str)-1; i++ {
            if str[i]-str[i+1] != minus {
                return false
            }
        }
    }
    return true
}
