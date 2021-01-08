package main

import (
    "bufio"
    "os"
    "strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

type Point struct {
    Height  int
    Visited bool
}

var (
    Map    [][]Point
    N, max int
)

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N = nextInt()
    Map = make([][]Point, N, N)
    for y := 0; y < N; y++ {
        Map[y] = make([]Point, N, N)
        for x := 0; x < N; x++ {
            h := nextInt()
            max = Max(max, h)
            Map[y][x].Height = h
            Map[y][x].Visited = false
        }
    }
    wt.WriteString(strconv.Itoa(GetArea()))
}

type Queue struct {
    Values [][2]int
    Size   int
}

func (queue *Queue) Push(point [2]int) {
    (*queue).Values = append((*queue).Values, point)
    (*queue).Size++
}
func (queue *Queue) Pop() (point [2]int, ok bool) {
    if l := len((*queue).Values); (*queue).Size > 0 {
        point = (*queue).Values[l-1]
        (*queue).Values = (*queue).Values[:l-1]
        (*queue).Size--
        ok = true
    }
    return
}

var mvs = [4][2]int{
    {1, 0}, {-1, 0}, {0, -1}, {0, 1},
}

func BFS() (count int) {
    queue := Queue{}
    for y := range Map {
        for x := range Map[y] {
            if !Map[y][x].Visited {
                queue.Push([2]int{y, x})
                Map[y][x].Visited = true
                count++
                for queue.Size > 0 {
                    point, _ := queue.Pop()
                    for _, mv := range mvs {
                        ym, xm := point[0]+mv[0], point[1]+mv[1]
                        if 0 <= ym && ym < N && 0 <= xm && xm < N && !Map[ym][xm].Visited {
                            queue.Push([2]int{ym, xm})
                            Map[ym][xm].Visited = true
                        }
                    }
                }
            }
        }
    }
    return
}

func Flood(height int) {
    for y := range Map {
        for x := range Map[y] {
            if Map[y][x].Height == height && !Map[y][x].Visited {
                Map[y][x].Visited = true
            } else if Map[y][x].Height > height {
                Map[y][x].Visited = false
            }
        }
    }
}
func GetArea() (area int) {
    area = 1
    for height := 1; height < max; height++ {
        Flood(height)
        area = Max(area, BFS())
    }
    return
}

func Max(a, b int) (max int) {
    max = a
    if a < b {
        max = b
    }
    return
}
