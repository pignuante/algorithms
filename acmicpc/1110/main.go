package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    rd := bufio.NewReader(os.Stdin)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    var N string
    fmt.Fscanf(rd, "%s", &N)

    if len(N) == 1 {
        N = "0" + N
    }

    t := N
    count := 0
    for {
        r := sum(t)

        count += 1
        l := len(r)
        t = string(t[1]) + string(r[l-1])
        if t == N {
            break
        }
    }
    wt.WriteString(strconv.Itoa(count))
}

func sum(value string) (result string) {
    a, _ := strconv.Atoi(string(value[0]))
    b, _ := strconv.Atoi(string(value[1]))
    result = strconv.Itoa(a + b)
    return
}
