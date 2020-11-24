package main

import (
    "bufio"
    "os"
    "strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer wt.Flush()
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    T, _ := strconv.Atoi(sc.Text())
    for i := 0; i < T; i++ {
        sc.Scan()
        t, _ := strconv.Atoi(sc.Text())
        count(t)
    }
}

func count(num int) {
    zero := []int{1, 0}
    one := []int{0, 1}
    if num >= 2 {
        for i := 2; i <= num; i++ {
            zero = append(zero, zero[i-1]+zero[i-2])
            one = append(one, one[i-1]+one[i-2])
        }
    }
    wt.WriteString(strconv.Itoa(zero[num]) + " " + strconv.Itoa(one[num]) + "\n")
}
