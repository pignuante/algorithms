package main

import (
    "bufio"
    "os"
    "sort"
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

    for i := 0; i < N; i++ {
        sc.Scan()
        trees[i], _ = strconv.Atoi(sc.Text())
    }
    result := sort.Search(1<<32, func(i int) bool {
        sum := 0
        for _, tree := range trees {
            if tree > i {
                sum += tree - i
            }
        }
        return sum < M
    })
    wt.WriteString(strconv.Itoa(result - 1))
}
