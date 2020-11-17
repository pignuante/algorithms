package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {

    // sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanWords)
}

func main() {
    var min, max int
    min = 1000000
    max = -1000000

    writer := bufio.NewWriter(os.Stdout)
    input := nextInt()
    defer writer.Flush()

    for i := 0; i < input; i++ {
        v := nextInt()
        if min > v {
            min = v
        }
        if max < v {
            max = v
        }
    }
    writer.WriteString(strconv.Itoa(min) + " " + strconv.Itoa(max))
}

func nextInt() (r int) {
    if sc.Scan() {
        r, _ = strconv.Atoi(sc.Text())
    }
    return
}
