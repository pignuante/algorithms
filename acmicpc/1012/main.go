package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var mvs [][]int = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
    wt := bufio.NewWriter(os.Stdout)

    defer wt.Flush()

    sc.Scan()
    T, _ := strconv.Atoi(sc.Text())

    for i := 0; i < T; i++ {
        sc.Scan()
        MNK := strings.Fields(sc.Text())
        M, _ := strconv.Atoi(MNK[0])
        N, _ := strconv.Atoi(MNK[1])
        K, _ := strconv.Atoi(MNK[2])
        field := readField(M, N, K)
        var results []int
        for y := 0; y < N; y++ {
            for x := 0; x < M; x++ {
                if field[y][x] == 1 {
                    if t := BFS(field, x, y); t > 0 {
                        results = append(results, t)
                    }
                }
            }
        }
        wt.WriteString(strconv.Itoa(len(results)) + "\n")
    }
}

func BFS(field [][]int, x, y int) (count int) {
    xy := []int{x, y}
    queue := [][]int{xy}
    sizeY, sizeX := len(field), len(field[0])
    count = 1
    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]
        field[y][x] = 0
        for _, mv := range mvs {
            xt, yt := x+mv[0], y+mv[1]
            // fmt.Println(x,y,xt, yt, queue)
            if 0 <= xt && xt < sizeX && 0 <= yt && yt < sizeY && field[yt][xt] == 1 {
                field[yt][xt] = 0
                queue = append(queue, []int{xt, yt})
                count++
            }
        }
        queue = queue[1:]
    }
    return
}

func readField(M, N, K int) (field [][]int) {
    field = make([][]int, N)
    for k := range field {
        field[k] = make([]int, M)
    }
    for ; K > 0; K-- {
        sc.Scan()
        t := strings.Fields(sc.Text())
        x, _ := strconv.Atoi(t[0])
        y, _ := strconv.Atoi(t[1])
        field[y][x] = 1
    }
    return
}
