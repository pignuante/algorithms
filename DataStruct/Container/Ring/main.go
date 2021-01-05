package main

import (
    "container/ring"
    "fmt"
    "time"
)

func main() {
    dt := []string{"ONE", "TWO", "THREE", "FOUR", "FIVE"}

    r := ring.New(len(dt))
    for i := 0; i < r.Len(); i++ {
        r.Value = dt[i]
        r = r.Next()
    }

    r.Do(func(x interface{}) {
        fmt.Println(x)
    })

    fmt.Println("------")

    for range time.Tick(time.Second * 1) {
        fmt.Println(r.Value)
        r = r.Next()
    }
}
