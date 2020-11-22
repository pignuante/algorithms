package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
    "strings"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    line, _, _ := rd.ReadLine()
    temp := strings.Split(string(line), " ")
    N, _ := strconv.Atoi(temp[0])
    M, _ := strconv.Atoi(temp[1])

    cards := make([]int, N)
    line, _, _ = rd.ReadLine()
    nums := strings.Split(string(line), " ")
    for i := 0; i < N; i++ {
        temp, _ := strconv.Atoi(nums[i])
        cards[i] = temp
    }

    result := blackJack(cards, M)
    wt.WriteString(strconv.Itoa(result))
}

func blackJack(cards []int, M int) int {
    sort.Ints(cards)
    max := 0
    l := len(cards)

    for i := 0; i < l-2; i++ {
        for j := i + 1; j < l-1; j++ {
            for k := j + 1; k < l; k++ {
                temp := cards[i] + cards[j] + cards[k]
                if temp > max && temp <= M {
                    max = temp
                } else {
                    continue
                }
            }
        }
    }
    return max
}
