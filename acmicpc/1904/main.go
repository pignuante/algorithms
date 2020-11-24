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

    var mod = 15746
    N, _ := strconv.Atoi(sc.Text())
    a, b := 1, 1
    for ; N > 0; N-- {
        a, b = b, (a+b)%mod
    }
    wt.WriteString(strconv.Itoa(a))
}

// func makeDigits(num int) int {
// num1, num2 := 1, 1
// for ; num > 0; num-- {
// num1, num2 = num2, (num1+num2)%mod
// }
// return num1
// }
