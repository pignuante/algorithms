package main

import (
    "bufio"
    "os"
    "sort"
    "strconv"
)

var wt *bufio.Writer = bufio.NewWriter(os.Stdout)
var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

type Pair struct {
    Key   byte
    Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N := nextInt()
    nums := make([][]byte, N, N)
    express := make(map[byte]int)

    for i := 0; i < N; i++ {
        sc.Scan()
        nums[i] = Reverse(sc.Bytes())
        digits := 1
        for _, v := range nums[i] {
            express[v] += digits
            digits *= 10
        }
    }
    pairList := make(PairList, len(express), len(express))
    i := 0
    for k, v := range express {
        pairList[i] = Pair{k, v}
        i++
    }
    sort.Sort(pairList)
    result := 0
    for k, v := range pairList {
        result += (9 - k) * v.Value
    }
    wt.WriteString(strconv.Itoa(result))

}

func Reverse(numbers []byte) []byte {
    for i := 0; i < len(numbers)/2; i++ {
        j := len(numbers) - i - 1
        numbers[i], numbers[j] = numbers[j], numbers[i]
    }
    return numbers
}
