package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var mod = 15746

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    fmt.Println(makeDigits(N))
}

func makeDigits(num int) int {
    nums := []int{0, 1, 2}
    if num >= 3 {
        for i := 3; i < num+1; i++ {
            nums = append(nums, (nums[i-2]+nums[i-1])%mod)
        }
    }
    return nums[num]
}
