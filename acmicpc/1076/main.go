package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    reg := map[string][2]int {
        "black":{0, 1}, "brown":{1, 10}, "red":{2, 100}, "orange":{3, 1000},
        "yellow":{4, 10000}, "green":{5, 100000}, "blue":{6, 1000000},
        "violet":{7, 10000000}, "grey":{8, 100000000}, "white":{9, 1000000000},
    }
    sum := 10
    sc.Scan()
    sum *= reg[sc.Text()][0]
    sc.Scan()
    sum += reg[sc.Text()][0]
    sc.Scan()
    sum *= reg[sc.Text()][1]
    wt.WriteString(strconv.Itoa(sum))
}
