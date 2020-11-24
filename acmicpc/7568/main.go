package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

type bulk struct {
    weight int
    height int
    rank   int
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    persons := make([]bulk, N)

    for k := range persons {
        sc.Scan()
        text := strings.Fields(sc.Text())
        weight, _ := strconv.Atoi(text[0])
        height, _ := strconv.Atoi(text[1])
        persons[k].weight = weight
        persons[k].height = height
    }

    orderBulk(persons)
    for _, v := range persons {
        wt.WriteString(strconv.Itoa(v.rank) + " ")
    }

}

func orderBulk(persons []bulk) {
    for k, v := range persons {
        count := 1
        for _, vv := range persons {
            if v.weight < vv.weight && v.height < vv.height {
                count++
            }
        }
        persons[k].rank = count
    }
}
