package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    times := make([]int, N, N)
    sc.Scan()
    nums := strings.Fields(sc.Text())
    for k, v := range nums {
        times[k], _ = strconv.Atoi(v)
    }
    sort.Ints(times)
    r := 0
    before := 0
    for _, v := range times {
        r += before + v
        before += v
    }
    // wt.WriteString(strconv.Itoa(sum(times)))
    wt.WriteString(strconv.Itoa(r))
}

func sum(times []int) (r int) {
    before := 0
    for _, v := range times {
        r += before + v
        before += v
    }
    return
}
