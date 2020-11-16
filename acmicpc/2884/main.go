package main

import "fmt"

func main() {
    var hour, minute int

    fmt.Scanf("%d %d", &hour, &minute)

    if minute < 45 {
        hour = (hour-1+24) % 24
        minute = 60-(45-minute)
    } else {
        minute -= 45
    }
    fmt.Println(hour, minute)

}
