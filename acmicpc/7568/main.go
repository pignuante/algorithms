package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    persons := make([][2]int, N)

    for k := range persons {
        sc.Scan()
        text := strings.Fields(sc.Text())
        weight, _ := strconv.Atoi(text[0])
        height, _ := strconv.Atoi(text[1])
        persons[k][0] = weight
        persons[k][1] = height
    }

    result := orderBulk(persons)
    for _, v := range result {
        wt.WriteString(strconv.Itoa(v) + " ")
    }

}

func orderBulk(persons [][2]int) []int {
    l := len(persons)
    rank := make([]int, l)
    for k, v := range persons {
        count := 1
        for _, vv := range persons {
            if v[0] < vv[0] && v[1] < vv[1] {
                count++
            }
        }
        rank[k] = count
    }
    return rank
}
