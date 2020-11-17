package main

import (
    "os"
    "bufio"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)
func init() {
    sc.Split(bufio.ScanWords)
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    num := 9+1

    max := -1
    idx := -1
    for i := 1; i < num; i++ {
        t := nextInt()
        if t > max {
            max = t
            idx = i
        }
    }
    wt.WriteString(strconv.Itoa(max) + "\n")
    wt.WriteString(strconv.Itoa(idx))

}

func nextInt() (r int){
    sc.Scan()
    r, _ = strconv.Atoi(sc.Text())
    return
}
