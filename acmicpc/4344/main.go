package main

import (
    "bufio"
    "fmt"
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
        scores := make([]float64, N)
        var sum float64 = 0.0
        for n := 0; n < N; n++ {
            scores[n], _ = strconv.ParseFloat(nextString(), 64)
            sum += scores[n]
        }
        mean := sum / float64(N)
        count := 0.0
        for n := 0; n < N; n++ {
            if scores[n] > mean {
                count++
            }
        }
        fmt.Fprintf(wt, "%0.3f%%\n", count/float64(N)*100)
    }
}

func nextString() (r string) {
    if sc.Scan() {
        r = sc.Text()
    }

    return
}
