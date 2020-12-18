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

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    board := make([][]bool, N, N)
    for k := range board {
        board[k] = make([]bool, N, N)
    }
    var check [3][40]bool
    count := 0
    Queen(&board, &check, &count, 0, N)
    wt.WriteString(strconv.Itoa(count))
}

func Queen(board *[][]bool, check *[3][40]bool, count *int, idx, N int) {
    if idx == N {
        *count++
    }
    for i := 0; i < N; i++ {
        if (*check)[0][i] || (*check)[1][idx+i] || (*check)[2][N+idx-i] {
            continue
        }
        (*check)[0][i], (*check)[1][idx+i], (*check)[2][N+idx-i] = true, true, true
        Queen(board, check, count, idx+1, N)
        (*check)[0][i], (*check)[1][idx+i], (*check)[2][N+idx-i] = false, false, false
    }
}
