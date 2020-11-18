package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func init() {
    sc.Split(bufio.ScanWords)
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    C, _ := strconv.Atoi(nextString())
    fmt.Println(C)

    for c := 0; c < C; c++ {
        N, _ := strconv.Atoi(nextString())
        scores := make([]int, N)
        var sum int = 0.0
        for n := 0; n < N; n++ {
            scores[n], _ = strconv.Atoi(nextString())
            sum += scores[n]
        }
        // mean := floatRound(float64(sum) / float64(N), 0.005)
        mean := float64(sum) / float64(N)
        count := 0.0
        for n := 0; n < N; n++ {
            if float64(scores[n]) > mean {
                count++
            }
        }
        fmt.Fprintf(wt, "%0.3f%%\n", floatRound(count/float64(N)*100, 3))
    }
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func floatRound(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num*output)) / output
}

func nextString() (r string) {
    if sc.Scan() {
        r = sc.Text()
    }

    return
}
