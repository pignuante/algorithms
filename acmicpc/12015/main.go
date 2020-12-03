package main

import (
    "bufio"
    "os"
    "strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    r = 0
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}
func main() {
    sc.Split(bufio.ScanWords)
    wt := bufio.NewWriter(os.Stdout)

    defer wt.Flush()
    A := nextInt()
    var vector []int
    vector = append(vector, nextInt())
    for i := 1; i < A; i++ {
        lowerBounder(&vector, nextInt())
    }
    wt.WriteString(strconv.Itoa(len(vector)))
}

func lowerBounder(vector *[]int, num int) {
    if (*vector)[len(*vector)-1] < num {
        *vector = append(*vector, num)
    } else {
        for i := 0; i < len(*vector); i++ {
            if (*vector)[i] >= num {
                (*vector)[i] = num
                break
            }
        }
    }
}
