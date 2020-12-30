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

var (
    N, M    int
    count   = -1
    Map     [][]byte
    viruses [][2]int
    stack   []int
    mvs     = [][2]int{
        {1, 0}, {-1, 0}, {0, -1}, {0, 1},
    }
)

const Walls = 3

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, M = nextInt(), nextInt()
    Map = make([][]byte, N, N)
    for y := 0; y < N; y++ {
        Map[y] = make([]byte, M, M)
        for x := 0; x < M; x++ {
            sc.Scan()
            status := sc.Bytes()[0]
            Map[y][x] = status
            if status == '2' {
                viruses = append(viruses, [2]int{y, x})
            }
        }
    }

    MakeWalls()
    wt.WriteString(strconv.Itoa(count))
}
func Spread() {
    AfterMap := make([][]byte, N, N)
    for y := range Map {
        AfterMap[y] = make([]byte, M, M)
        copy(AfterMap[y], Map[y])
    }

    var queue [][2]int
    queue = append(queue, viruses...)
    for len(queue) > 0 {
        point := queue[0]
        queue = queue[1:]
        for _, mv := range mvs {
            ym, xm := point[0]+mv[0], point[1]+mv[1]
            if 0 <= ym && ym < N && 0 <= xm && xm < M {
                if AfterMap[ym][xm] == '0' {
                    AfterMap[ym][xm] = '2'
                    queue = append(queue, [2]int{ym, xm})
                }
            }
        }
    }
    count = Max(CountSafe(AfterMap), count)
}
func CountSafe(afterMap [][]byte) (c int) {
    c = 0
    for _, y := range afterMap {
        for _, v := range y {
            if v == '0' {
                c++
            }
        }
    }
    return
}

func MakeWalls() {
    if len(stack) == 3 {
        Spread()
        return
    }

    for y := 0; y < N; y++ {
        for x := 0; x < M; x++ {
            if value := y*M + x; Map[y][x] == '0' { // 빈공간이고
                if l := len(stack); l == 0 || stack[l-1] < value { // 이전에 방문한곳이 아니라면
                    stack = append(stack, value)
                    Map[y][x] = '1'
                    MakeWalls()
                    Map[y][x] = '0'
                    stack = stack[:len(stack)-1]
                }
            }
        }
    }
}
func Max(a, b int) (max int) {
    max = a
    if a < b {
        max = b
    }
    return
}
