package main

import (
    "bufio"
    "os"
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

    N, K := nextInt(), nextInt()
    list := make([]bool, N+1, N+1)
    list[0], list[1] = true, true

    count, flag := 0, false
    result := 2
    for i := 2; i <= N; i++ {
        for j := i; j <= N; j += i {
            if !list[j] {
                list[j] = true
                count++
            }
            if count == K {
                result = j
                flag = true
                break
            }
        }
        if flag {
            break
        }
    }

    wt.WriteString(strconv.Itoa(result))
}
