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
    var nums []int

    result := getReminder(nums)
    printResult(result)

}

func printResult(nums []int) {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    wt.WriteString(strconv.Itoa(len(nums)))

}

func getReminder(nums []int) []int {
    num := 10
    std := 42

    for i := 0; i < num; i++ {
        t := nextInt()

        var flag bool = true
        for _, v := range nums {
            if v == t%std {
                flag = !flag
            }
        }
        if flag {
            nums = append(nums, t%42)
        }
    }
    return nums
}

func nextInt() (r int) {
    if sc.Scan() {
        r, _ = strconv.Atoi(sc.Text())
    }

    return
}
