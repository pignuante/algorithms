package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r = 0
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}
func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    wt.WriteString(strconv.Itoa(Sugar(N)))
}

func Sugar(kg int) (bags int) {
    five := kg / 5
    three := 0
    for temp := kg % 5; five >= 0; {
        if temp%3 == 0 {
            three = temp / 3
            temp %= 3
            break
        }
        five -= 1
        temp += 5
    }
    bags = five + three
    return
}
