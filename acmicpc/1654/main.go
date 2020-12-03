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
    K, _ := strconv.Atoi(sc.Text())
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    lines := make([]int, K)

    end := 1
    for i := 0; i < K; i++ {
        sc.Scan()
        n, _ := strconv.Atoi(sc.Text())
        lines[i] = n
        if end < n {
            end = n
        }
    }
    start := 1
    mid := 0
    result := 0
    for start <=  end {
        mid = (start+end)>>1
        sumOfLines := 0
        for _, v := range lines {
            sumOfLines += v/mid
        }
        if sumOfLines < N {
            end = mid - 1
        } else if sumOfLines >= N {
            result = mid
            start = mid + 1
        }
    }
    wt.WriteString(strconv.Itoa(result))
}
