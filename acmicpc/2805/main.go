package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanWords)

    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    sc.Scan()
    M, _ := strconv.Atoi(sc.Text())

    trees := make([]int, N, N)

    start, end := 0, 0
    for i := 0; i < N; i++ {
        sc.Scan()
        trees[i], _ = strconv.Atoi(sc.Text())
        if end < trees[i] {
            end = trees[i]
        }
    }

    mid, result := 0, 0
    for start <= end {
        mid = (start + end) >> 1
        sumOfTree := 0
        for _, tree := range trees {
            if tree > mid {
                sumOfTree += (tree - mid)
            }
        }
        if sumOfTree < M {
            end = mid - 1
        } else if sumOfTree >= M {
            result = mid
            start = mid + 1
        }
    }
    wt.WriteString(strconv.Itoa(result))
}
