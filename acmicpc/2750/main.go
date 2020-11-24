package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)

    defer wt.Flush()
    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    nums := make([]int, N, N)

    for ; N > 0; N-- {
        sc.Scan()
        n, _ := strconv.Atoi(sc.Text())
        nums[N-1] = n
    }
    bubbleSort(nums)
    for _, v := range nums {
        wt.WriteString(strconv.Itoa(v) + "\n")
    }
}

func bubbleSort(nums []int) {
    for i := range nums {
        for j := range nums[:len(nums)-(i+1)] {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
}
