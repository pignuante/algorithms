package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	r, f := 0, 1
	for _, c := range sc.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

type Nums struct {
	Num, Order int
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	N := nextInt()
	nums := make([]int, N, N)
	temp := make([]Nums, N, N)
	for i := 0; i < N; i++ {
		num := nextInt()
		nums[i] = num
		temp[i] = Nums{num, i}
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].Num < temp[j].Num
	})
	result := make(map[int]int)
	idx := 0
	for _, value := range temp {
		if _, isExist := result[value.Num]; !isExist {
			result[value.Num] = idx
			idx++
		}
	}
	for _, num := range nums {
		wt.WriteString(strconv.Itoa(result[num]))
		wt.WriteByte('\n')
	}
}
