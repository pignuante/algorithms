package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    nums := 9
    dwarfs := make([]int, nums, nums)
    sum := 0
    for i := 0; i < nums; i++ {
        dwarfs[i] = nextInt()
        sum += dwarfs[i]
    }
    sort.Ints(dwarfs)
    for i := 0; i < nums; i++ {
        for j := i + 1; j < nums; j++ {
            if sum-dwarfs[i]-dwarfs[j] == 100 {
                for k := 0; k < nums; k++ {
                    if k != i && k != j {
                        wt.WriteString(strconv.Itoa(dwarfs[k]))
                        wt.WriteByte('\n')
                    }
                }
                return
            }
        }
    }
}
