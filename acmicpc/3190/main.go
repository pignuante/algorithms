package main

import (
    "bufio"
    "container/list"
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
    Apple, Visited bool
}
type Command struct {
    Turn int
    Cmd  byte
}

var (
    N, K, L  int
    Map      [][]Point
    Snake    list.List
    Commands []Command
    Mvs      = [][2]int{
        {0, 1}, {1, 0}, {0, -1}, {-1, 0},
    }
)

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N = nextInt()
    Map = make([][]Point, N+1, N+1)
    for y := 0; y <= N; y++ {
        Map[y] = make([]Point, N+1, N+1)
    }
    K = nextInt()
    for i := 0; i < K; i++ {
        y, x := nextInt(), nextInt()
        Map[y][x].Apple = true
    }
    L = nextInt()
    Commands = make([]Command, L, L)
    for i := 0; i < L; i++ {
        Commands[i].Turn = nextInt()
        sc.Scan()
        Commands[i].Cmd = sc.Bytes()[0]
    }
    wt.WriteString(strconv.Itoa(Dummy()))
}

func Dummy() (turn int) {
    y, x, c := 1, 1, 0
    Snake.PushFront([2]int{y, x})
    mv, turn := 0, 0
    for {
        turn++
        y = (y + Mvs[mv][0])
        x = (x + Mvs[mv][1])
        if y < 1 || y >= N+1 || x < 1 || x >= N+1 || Map[y][x].Visited {
            break
        }
        Snake.PushFront([2]int{y, x})
        Map[y][x].Visited = true
        if !Map[y][x].Apple {
            t := Snake.Remove(Snake.Back()).([2]int)
            Map[t[0]][t[1]].Visited = false

        } else {
            Map[y][x].Apple = false
        }
        if c < L && turn == Commands[c].Turn {
            switch Commands[c].Cmd {
            case 'D':
                mv = (mv + 1) % 4
            case 'L':
                mv = (mv - 1) % 4
            }
            c++
        }
    }
    return
}
