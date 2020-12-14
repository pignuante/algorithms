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

type country struct {
    population       int
    visited, changed bool
}

var mvs = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, L, R := nextInt(), nextInt(), nextInt()
    countries := make([][]country, N, N)
    for y := 0; y < N; y++ {
        countries[y] = make([]country, N, N)
        for x := 0; x < N; x++ {
            countries[y][x].population = nextInt()
        }
    }

    wt.WriteString(strconv.Itoa(MovePop(&countries, L, R) - 1))
}

func MovingPop(countries *[][]country, pop int) {
    l := len(*countries)
    for y := 0; y < l; y++ {
        for x := 0; x < l; x++ {
            if (*countries)[y][x].changed {
                (*countries)[y][x].changed = false
                (*countries)[y][x].population = pop
            }
        }
    }
}
func InitCountries(countries *[][]country) {
    for y := range *countries {
        for x := range *countries {
            (*countries)[y][x].visited = false
            (*countries)[y][x].changed = false
        }
    }
}
func MovePop(countries *[][]country, L, R int) (moved int) {
    flag, l, moved := true, len(*countries), 0

    for flag {
        flag = false
        moved++
        for y := 0; y < l; y++ {
            for x := 0; x < l; x++ {
                queue := [][2]int{{y, x}}
                pop, count := 0, 0
                for len(queue) > 0 {
                    y, x := queue[0][0], queue[0][1]
                    queue = queue[1:]
                    for _, mv := range mvs {
                        ym, xm := y+mv[0], x+mv[1]
                        if 0 <= ym && ym < l && 0 <= xm && xm < l && !(*countries)[ym][xm].visited {
                            dis := Abs((*countries)[ym][xm].population - (*countries)[y][x].population)
                            if L <= dis && dis <= R {
                                flag = true
                                pop += (*countries)[ym][xm].population
                                count++
                                (*countries)[ym][xm].visited = true
                                (*countries)[ym][xm].changed = true
                                queue = append(queue, [2]int{ym, xm})
                            }
                        }
                    }
                }
                if count > 0 {
                    MovingPop(countries, pop/count)
                }
            }
        }
        InitCountries(countries)
    }
    return
}

func Abs(num int) (r int) {
    r = num
    if r < 0 {
        r = -r
    }
    return
}
