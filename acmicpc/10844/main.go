package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var mod int = 1000000000

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    stairs := make([][10]int, N)
    for i := 1; i < 10; i++ {
        stairs[0][i] = 1
    }
    GetStairsNum(&stairs)
    wt.WriteString(strconv.Itoa(Sum(stairs[N-1]) % mod))
}

func Sum(nums [10]int) (sum int) {
    for _, v := range nums {
        sum += v
    }
    return
}
func GetStairsNum(stairs *[][10]int) {
    for y := 1; y < len(*stairs); y++ {
        for x := 0; x < 10; x++ {
            if x == 0 {
                (*stairs)[y][x] = (*stairs)[y-1][1] % mod
            } else if x == 9 {
                (*stairs)[y][x] = (*stairs)[y-1][8] % mod
            } else {
                (*stairs)[y][x] = ((*stairs)[y-1][x-1] + (*stairs)[y-1][x+1]) % mod
            }
        }
    }
}
