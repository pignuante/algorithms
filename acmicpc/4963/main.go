package main

import (
    "bufio"
    "os"
    "strconv"
)

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
    y, x int
}

var mvs = [9][2]int{
    {0, 0},
    {1, 0}, {-1, 0}, {0, -1}, {0, 1},
    {1, -1}, {1, 1}, {-1, -1}, {-1, 1},
}
var w, h int

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    for {
        w, h = nextInt(), nextInt()
        if w == 0 && h == 0 {
            break
        }
        Map := make([][]int, h, h)
        visited := make([][]bool, h, h)
        for y := 0; y < h; y++ {
            for x := 0; x < w; x++ {
                Map[y] = append(Map[y], nextInt())
            }
            visited[y] = make([]bool, w, w)
        }
        count := 0
        for y := 0; y < h; y++ {
            for x := 0; x < w; x++ {
                if FindMap(&Map, &visited, y, x) {
                    count++
                }
            }
        }
        wt.WriteString(strconv.Itoa(count))
        wt.WriteByte('\n')
    }
}

func FindMap(Map *[][]int, visited *[][]bool, y, x int) (flag bool) {
    queue := []Point{{y: y, x: x}}
    for len(queue) > 0 {
        point := queue[0]
        queue = queue[1:]
        for _, mv := range mvs {
            ym, xm := point.y+mv[0], point.x+mv[1]
            if 0 <= ym && ym < h && 0 <= xm && xm < w {
                if (*Map)[ym][xm] == 1 && !(*visited)[ym][xm] {
                    (*visited)[ym][xm] = true
                    queue = append(queue, Point{y: ym, x: xm})
                    flag = true
                }
            }
        }
    }
    return flag
}
