package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Member struct {
    Age   int
    Name  string
    Order int
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    member := make([]Member, N, N)
    for i := 0; i < N; i++ {
        sc.Scan()
        t := strings.Fields(sc.Text())
        member[i].Age, _ = strconv.Atoi(t[0])
        member[i].Name = t[1]
        member[i].Order = i
    }
    sort.Slice(member, func(i, j int) bool {
        if member[i].Age == member[j].Age {
            return member[i].Order < member[j].Order
        }
        return member[i].Age < member[j].Age
    })
    for _, v := range member {
        wt.WriteString(strconv.Itoa(v.Age) + " " + v.Name + "\n")
    }
}
