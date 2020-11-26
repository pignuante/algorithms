package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

type Computer struct {
    networks []int
    status   bool
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    sc.Scan()
    cs, _ := strconv.Atoi(sc.Text())
    computers := make([]Computer, cs+1, cs+1)

    sc.Scan()
    nets, _ := strconv.Atoi(sc.Text())
    for i := 0; i < nets; i++ {
        sc.Scan()
        nums := strings.Fields(sc.Text())
        a, _ := strconv.Atoi(nums[0])
        b, _ := strconv.Atoi(nums[1])
        computers[a].networks = append(computers[a].networks, b)
        computers[b].networks = append(computers[b].networks, a)
    }
    stack := 1
    DFS(computers, 1, &stack)
    wt.WriteString(strconv.Itoa(stack-1))
}
func DFS(computers []Computer, idx int, stack *int) {
    computers[idx].status = true
    for _, network := range computers[idx].networks {
        if computers[network].status == false {
            *stack++
            DFS(computers, network, stack)
        }
    }
}
