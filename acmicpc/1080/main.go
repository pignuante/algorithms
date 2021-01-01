package main

import (
    "bufio"
    "os"
    "strconv"
    "fmt"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c-'0')
    }
    return
}
var (
    N, M, count int
    compare [][]bool
    arr [2][][]bool
    flag bool
)
func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N, M = nextInt(), nextInt()
    for i := 0; i < 2; i++ {
        arr[i] = make([][]bool, N, N)
        for y := 0; y < N; y++ {
            arr[i][y] = make([]bool, M, M)
            sc.Scan()
            for x, v := range sc.Bytes() {
                if v == '0' {
                    arr[i][y][x] = false
                } else {
                    arr[i][y][x] = true
                }
            }
        }
    }
    for y := 0; y < N-2; y++ {
        for x := 0; x < M-2; x++ {
            if arr[0][y][x] != arr[1][y][x] {
                flag = true
                Flip(y, x)
                count++
            }
        }
    }
    for i := 0; i < 2; i++ {
        for _, v := range arr[i] {
            fmt.Println(v)
        }
        fmt.Println()
    }
    if Compare(){
        wt.WriteString(strconv.Itoa(count))
    } else {
        wt.WriteString("-1")
    }
}


func Compare() (bool){
    for y := 0; y < N; y++ {
        for x := 0; x < M; x++ {
            if arr[0][y][x] != arr[1][y][x] {
                return false
            }
        }
    }
    return true
}

func Flip(row, column int) {
    for y := row; y < row + 3; y++ {
        for x := column; x < column + 3; x++ {
            arr[0][y][x] = !arr[0][y][x]
        }
    }
}
