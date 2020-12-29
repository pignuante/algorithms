package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    ans, MOD := 0, 20000303
    N, _ := reader.ReadString('\n')
    for _, c := range N[:len(N)-1] {
        ans = (ans*10 + int(c-'0')) % MOD
    }
    fmt.Println(ans)
}
