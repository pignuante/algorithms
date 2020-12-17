package main

import "fmt"

func main() {
    var index int
    fmt.Scanf("%d", &index)

    var sections [][]string
    sections = append(sections, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
    for i := 0; i < 9; i++ {
        section := []string{}
        for j := len(sections); j <= 9; j++ {
            for _, n := range sections[len(sections)-1] {
                if fmt.Sprintf("%d", j)[0] > n[0] {
                    section = append(section, fmt.Sprintf("%d%s", j, n))
                }
            }
        }
        sections = append(sections, section)
    }
    var result []string
    for _, eachSection := range sections {
        result = append(result, eachSection...)
    }
    if index < 1023 {
        fmt.Println(result[index])
    } else {
        fmt.Println(-1)
    }
}
