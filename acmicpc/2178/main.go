package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

type Point struct {
    status, dist int
    visited            bool
}

var mvs [][]int = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
    wt := bufio.NewWriter(os.Stdout)

    defer wt.Flush()

    sc.Scan()
    NM := strings.Fields(sc.Text())
    N, _ := strconv.Atoi(NM[0])
    M, _ := strconv.Atoi(NM[1])

    maze := readMaze(N, M)
    BFS(maze, 0, 0, N, M)
    wt.WriteString(strconv.Itoa(maze[N-1][M-1].dist))
}

func BFS(maze [][]Point, y, x, N, M int) {
    maze[y][x].dist = 1
    queue := [][]int{{y, x}}
    for len(queue) > 0 {
        y, x = queue[0][0], queue[0][1]
        maze[y][x].visited = true
        for _, mv := range mvs {
            yt, xt := y+mv[0], x+mv[1]
            if 0 <= yt && yt < N && 0 <= xt && xt < M && maze[yt][xt].status == 1 && !maze[yt][xt].visited {
                maze[yt][xt].status = 3
                maze[yt][xt].dist = maze[y][x].dist + 1
                maze[yt][xt].visited = true
                queue = append(queue, []int{yt, xt})
            }
        }
        queue = queue[1:]
    }
}

func readMaze(Y, X int) (maze [][]Point) {
    maze = make([][]Point, Y, Y)
    for y := 0; y < Y; y++ {
        sc.Scan()
        t := sc.Text()
        X := make([]Point, len(t))
        for x, v := range t {
            X[x] = Point{status: int(v - '0'), dist: 0, visited: false}
        }
        maze[y] = X
    }
    return
}
