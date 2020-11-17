package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
    sc.Split(bufio.ScanWords)
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    N := nextInt()
    nums := make([]float64, N)
    max := -1.0
    for i := 0; i < N; i++ {
        t := float64(nextInt())
        nums[i] = t
        if max < t {
            max = t
        }
    }
    result := 0.0
    for _, v := range nums {
        result += (v / max) * 100
    }
    result /= float64(N)
    // r := floatToString(result)
    fmt.Fprintf(wt, "%0.4f", result)

}

// func floatToString(num float64) (r string) {
// num = Round(num, 0.0005)
// r = strconv.FormatFloat(num, 'g', 1, 64)
// return
// }

func Round(x, unit float64) float64 {
    return math.Round(x/unit) * unit
}

func nextInt() (num int) {
    if sc.Scan() {
        num, _ = strconv.Atoi(sc.Text())
    }
    return
}
