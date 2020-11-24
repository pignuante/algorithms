package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdout)
    defer wt.Flush()

    sc.Scan()
    NM := strings.Split(sc.Text(), " ")
    N, _ := strconv.Atoi(NM[0])
    M, _ := strconv.Atoi(NM[1])
    board := make([][]byte, N)
    for i := 0; i < N; i++ {
        sc.Scan()
        board[i] = sc.Bytes()
    }
    count := 99
    for y := 0; y < N-7; y++ {
        for x := 0; x < M-7; x++ {
            count = min(count, search(y, x, board))
        }
    }
    wt.WriteString(strconv.Itoa(count))
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func search(startY, startX int, board [][]byte) int {
    colors := [2]byte{'W', 'B'}
    counts := [2]int{0, 0}
    for y := startY; y < startY+8; y++ {
        for x := startX; x < startX+8; x++ {
            if board[y][x] != colors[(y+x)%2] {
                counts[0]++
            }
            if board[y][x] != colors[((y+x)+1)%2] {
                counts[1]++
            }
        }
    }
    return min(counts[0], counts[1])
}
