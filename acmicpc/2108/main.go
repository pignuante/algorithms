package main

import (
    "bufio"
    "math"
    "os"
    "sort"
    "strconv"
)

func main() {
    rangeHigh := 4000
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    nums := make([]int, rangeHigh*2+1, rangeHigh*2+1)
    input := make([]int, N, N)

    count := []int{-4001, -1}
    sum, min, max := 0, 4001, -4001
    for i := 0; i < N; i++ {
        sc.Scan()
        n, _ := strconv.Atoi(sc.Text())
        sum += n
        if min > n {
            min = n
        }
        if max < n {
            max = n
        }
        nums[n+rangeHigh]++
        if nums[n+rangeHigh] > count[1] {
            count[0] = n
            count[1] = nums[n+rangeHigh]
        }
        input[i] = n
    }
    sort.Ints(input)
    wt.WriteString(strconv.Itoa(int(math.Round(float64(sum)/float64(N)))) + "\n")
    wt.WriteString(strconv.Itoa(input[len(input)/2]) + "\n")
    mod := findMod(nums, count)
    wt.WriteString(strconv.Itoa(mod) + "\n")
    abs := max - min
    if abs < 0 {
        abs = -abs
    }
    wt.WriteString(strconv.Itoa(abs))
}

func findMod(nums, count []int) (mod int) {
    max := count[1]
    counter := 0
    for k, v := range nums {
        if v == max {
            mod = k - 4000
            counter++
            if counter == 2 {
                break
            }
        }
    }
    return
}
