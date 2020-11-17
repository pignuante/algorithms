package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
    sc.Split(bufio.ScanWords)
}
func main() {
    num := 3
    abc := make([]int, 3)
    for i := 0; i < num; i++ {
        abc[i] = nextInt()
    }
    mul := abc[0] * abc[1] * abc[2]

    printResult(countInt(mul))

}

func printResult(result []int) {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()
    for _, v := range result {
        wt.WriteString(strconv.Itoa(v) + "\n")
    }
}

func countInt(mul int) (result []int) {

    strNum := strconv.Itoa(mul)

    result = make([]int, 10)
    for _, v := range strNum {
        t, _ := strconv.Atoi(string(v))
        result[t] += 1
    }

    return
}

func nextInt() (r int) {
    sc.Scan()
    r, _ = strconv.Atoi(sc.Text())

    return
}
