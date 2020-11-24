package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    rooms := make([][2]int, N)

    for i := 0; i < N; i++ {
        sc.Scan()
        t := strings.Fields(sc.Text())
        rooms[i][0], _ = strconv.Atoi(t[0])
        rooms[i][1], _ = strconv.Atoi(t[1])
    }
    Sort(rooms)
    wt.WriteString(strconv.Itoa(allocateRoom(rooms)))
}

func Sort(rooms [][2]int) {
    sort.Slice(rooms, func(i, j int) bool {
        if rooms[i][1] < rooms[j][1] {
            return true
        }
        if rooms[i][1] > rooms[j][1] {
            return false
        }
        return rooms[i][0] < rooms[j][0]
    })
}

func allocateRoom(rooms [][2]int) (count int) {
    end := 0
    for _, time := range rooms {
        if time[0] >= end {
            end = time[1]
            count++
        }
    }
    return
}
