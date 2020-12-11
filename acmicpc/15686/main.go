package main

import (
    "bufio"
    "os"
    // "strconv"
    "fmt"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c-'0')
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N, M := nextInt(), nextInt()
    var houes, chickens [][2]int
    for y := 1; y <= N; y++ {
        for x := 1; x <= N; x++ {
            if t := nextInt(); t == 1 {
                houes = append(houes, [2]int{y, x})
            } else if t == 2 {
                chickens = append(chickens, [2]int{y, x})
            }
        }
    }
    fmt.Println(M)
    fmt.Println(houes)
    fmt.Println(chickens)
    fmt.Println(GetAllChickeDistance(houes, chickens))
}

func GetAllChickeDistance(houses, chickens [][2]int) (chickenDistances []int) {
    chickenDistances = make([]int, len(houses), len(houses))
    for k, house := range  houses {
        chickenDistances[k] = GetChickenDistance(house, chickens)
    }
    return
}

func GetChickenDistance(houes [2]int, chickens [][2]int) (min int) {
    min = 1000
    for _, chicken := range chickens {
        if dist := Manhattan(houes, chicken); min > dist {
            min = dist
        }
    }
    return
}

func Abs(num int) (int) {
    if num < 0 {
        num = -num
    }
    return num
}

func Manhattan(houes, chicken [2]int) (int) {
    return Abs(chicken[1] - houes[1]) + Abs(chicken[0] - houes[0])
}
