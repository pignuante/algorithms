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

var N int
var mvs = [][2]int{
    {1, 0}, {-1, 0}, {0, -1}, {0, 1},
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N = nextInt()
    Board, visited := make([][]byte, N, N), make([][]bool, N, N)
    for y := 0; y < N; y++ {
        sc.Scan()
        Board[y] =  sc.Bytes()
        visited[y] = make([]bool, N, N)
    }

    count := 0
    for y := 0; y < N; y++ {
        for x := 0; x < N; x++ {
            if !visited[y][x] {
                FindColor(&Board, &visited, y, x)
                count++
            }
        }
    }
    wt.WriteString(strconv.Itoa(count)+" ")

    for y := 0; y < N; y++ {
        for x := 0; x < N; x++ {
            if Board[y][x] == 'R' {
                Board[y][x] = 'G'
            }
            visited[y][x] = false
        }
    }
    count = 0
    for y := 0; y < N; y++ {
        for x := 0; x < N; x++ {
            if !visited[y][x] {
                FindColor(&Board, &visited, y, x)
                count++
            }
        }
    }
    wt.WriteString(strconv.Itoa(count)+"\n")
}

func FindColor(Board *[][]byte, visited *[][]bool, y,x int) {
    (*visited)[y][x] = true
    for _, mv := range mvs {
        ym, xm := y+mv[0], x+mv[1]
        if 0 <= ym && ym < N && 0 <= xm && xm < N {
            if (*Board)[y][x] == (*Board)[ym][xm] && !(*visited)[ym][xm] {
                FindColor(Board, visited, ym, xm)
            }
        }
    }
}
