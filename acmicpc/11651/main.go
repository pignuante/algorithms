package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type nums struct {
	x, y int
}

func main() {
	wt := bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
	defer wt.Flush()

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	coordinate := make([]nums, N, N)
	for ; N > 0; N-- {
		sc.Scan()
		t := strings.Fields(sc.Text())
		coordinate[N-1].x, _ = strconv.Atoi(t[0])
		coordinate[N-1].y, _ = strconv.Atoi(t[1])
	}
	sort.Slice(coordinate, func(i, j int) bool {
		if coordinate[i].y == coordinate[j].y {
			return coordinate[i].x < coordinate[j].x
		}
		return coordinate[i].y < coordinate[j].y
	})
	for _, v := range coordinate {
		x := strconv.Itoa(v.x)
		y := strconv.Itoa(v.y)
		wt.WriteString(x + " " + y + "\n")
	}
}
