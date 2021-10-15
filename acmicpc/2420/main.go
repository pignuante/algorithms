package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	result := Absolute(N - M)
	fmt.Printf("%d", result)
}

func Absolute(number int) int {
	if number < 0 {
		number *= -1
	}
	return number
}
