package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

var mvs = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()

    sc.Scan()
    MN := strings.Fields(sc.Text())
    M, _ := strconv.Atoi(MN[0])
    N, _ := strconv.Atoi(MN[1])

    box := make([][]int, N, N)

    for y := 0; y < N; y++ {
        sc.Scan()
        X := strings.Fields(sc.Text())
        box[y] = make([]int, M, M)
        for x, v := range X {
            box[y][x], _ = strconv.Atoi(v)
        }
    }
    var queue [][2]int
    for y := 0; y < N; y++ {
        for x := 0; x < M; x++ {
            if box[y][x] == 1 {
                queue = append(queue, [2]int{x, y})
            }
        }
    }
    wt.WriteString(strconv.Itoa(BFS(box, queue, M, N)))
}

func BFS(box [][]int, queue [][2]int, X, Y int) int {
    days := -1
    for len(queue) > 0 {
        days++
        l := len(queue)
        for i := 0; i < l; i++ {
            x, y := queue[0][0], queue[0][1]
            queue = queue[1:]
            for _, mv := range mvs {
                xt, yt := x+mv[0], y+mv[1]
                if 0 <= xt && xt < X && 0 <= yt && yt < Y && box[yt][xt] == 0 {
                    box[yt][xt] = 1
                    queue = append(queue, [2]int{xt, yt})
                }
            }
        }
    }
    for _, v := range box {
        for _, k := range v {
            if k == 0 {
                return -1
            }
        }
    }
    return days
}
