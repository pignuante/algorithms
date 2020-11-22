package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    text := sc.Text()
    n, _ := strconv.Atoi(text)
    start := len(text) * 9

    if start > n {
        start = n
    }

    result := 0
    for i := n - start; i < n+1; i++ {
        num := strconv.Itoa(i)
        nums := makeNumber(num)
        if i+sum(nums) == n {
            result = i
            break
        }
    }

    fmt.Println(result)
}

func sum(num []int) (r int) {
    for _, v := range num {
        r += v
    }
    return
}

func makeNumber(num string) []int {
    l := len(num)
    r := make([]int, l)
    for i := 0; i < l; i++ {
        r[i] = int(num[i] - 48)
    }
    return r
}
