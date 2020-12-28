package main

import (
    "bufio"
    "os"
    "sort"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

var (
    stack   []byte
    visited []bool
    str     []byte
    L, C    int
)

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    L, C = nextInt(), nextInt()
    visited = make([]bool, C, C)
    str = make([]byte, C, C)
    for i := 0; i < C; i++ {
        sc.Scan()
        str[i] = sc.Bytes()[0]
    }
    sort.Slice(str, func(i, j int) bool {
        return str[i] < str[j]
    })
    Recursive(0, 0)
}

func Recursive(vowels, consontats int) {
    if len(stack) == L {
        if vowels >= 1 && consontats >= 2 {
            for _, v := range stack {
                wt.WriteByte(v)
            }
            wt.WriteByte('\n')
        }
        return
    }
    for i, v := range str {
        if visited[i] {
            continue
        }
        if l := len(stack); l == 0 || stack[l-1] <= str[i] {
            visited[i] = true
            stack = append(stack, v)
            if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
                Recursive(vowels+1, consontats)
            } else {
                Recursive(vowels, consontats+1)
            }
            stack = stack[:len(stack)-1]
            visited[i] = false
        }
    }
}
