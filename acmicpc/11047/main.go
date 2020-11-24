package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    sc.Scan()
    NK := strings.Fields(sc.Text())
    N, _ := strconv.Atoi(NK[0])
    K, _ := strconv.Atoi(NK[1])
    coins := make([]int, N, N)

    for ; N > 0; N-- {
        sc.Scan()
        coins[N-1], _ = strconv.Atoi(sc.Text())
    }
    wt.WriteString(strconv.Itoa(countCoins(coins, K)))
}

func countCoins(coins []int, K int) int {
    count := 0
    for _, coin := range coins {
        count += K / coin
        K %= coin
    }
    return count
}
