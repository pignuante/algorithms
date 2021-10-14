package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	d := bytes.Fields(b)
	n, _ := strconv.Atoi(string(d[0]))
	r := 0
	s, k := 0, 0
	for _, c := range d[2] {
		switch s {
		case 0:
			if c == 'I' {
				s = 1
			}
		case 1:
			if c == 'O' {
				s = 2
			} else {
				k = 0
			}
		case 2:
			if c == 'I' {
				k++
				if k >= n {
					r++
				}
				s = 1
			} else {
				s, k = 0, 0
			}
		}
	}
	fmt.Println(r)
}
