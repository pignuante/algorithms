package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
)

type Coord struct {
    x, y int
}

var mvs = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    aparts := make([][]int, N)
    for i := 0; i < N; i++ {
        sc.Scan()
        a := []rune(sc.Text())
        t := make([]int, N)
        for k, v := range a {
            t[k] = int(v) - '0'
        }
        aparts[i] = t
    }
    var apartComplex []int
    for y := 0; y < N; y++ {
        for x := 0; x < N; x++ {
            if aparts[y][x] == 1 {
                comp := BFS(aparts, x, y, N)
                if comp != 0 {
                    apartComplex = append(apartComplex, comp)
                }
            }
        }
    }
    wt.WriteString(strconv.Itoa(len(apartComplex)) + "\n")
    sort.Ints(apartComplex)
    for _, v := range apartComplex {
        wt.WriteString(strconv.Itoa(v) + "\n")
    }
}

func BFS(aparts [][]int, x, y int, N int) (count int) {
    coord := Coord{x: x, y: y}
    queue := []Coord{coord}
    count = 1
    for len(queue) > 0 {
        y := queue[0].y
        x := queue[0].x
        aparts[y][x] = 0
        for _, mv := range mvs {
            yt, xt := y+mv[1], x+mv[0]
            if 0 <= yt && yt < N && 0 <= xt && xt < N && aparts[yt][xt] == 1 {
                // fmt.Println(yt, xt, queue)
                aparts[yt][xt] = 0
                queue = append(queue, Coord{x: xt, y: yt})
                count++
            }
        }
        queue = queue[1:]
    }
    return
}
