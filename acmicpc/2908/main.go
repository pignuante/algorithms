package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    defer wt.Flush()

    rd := bufio.NewReader(os.Stdin)
    var nums [2]string
    fmt.Fscanf(rd, "%s %s", &nums[0], &nums[1])
    reverseNums(&nums)
    fmt.Fprintf(wt, "%d", findBiggestNum(nums))
}

func findBiggestNum(nums [2]string) int {
    a, _ := strconv.Atoi(nums[0])
    b, _ := strconv.Atoi(nums[1])
    if a > b {
        return a
    } else {
        return b
    }
}

func reverseNums(nums *[2]string) {
    for k, num := range nums {
        var reverse string = ""
        for _, vv := range num {
            reverse = string(vv) + reverse
        }
        nums[k] = reverse
    }
}
