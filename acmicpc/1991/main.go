package main

import (
    "bufio"
    "os"
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

type Node struct {
    Left, Right string
}

var Tree = make(map[string]Node)

func Preorder(node string) string {
    if node == "." {
        return ""
    }
    return node + Preorder(Tree[node].Left) + Preorder(Tree[node].Right)
}
func Inorder(node string) string {
    if node == "." {
        return ""
    }
    return Inorder(Tree[node].Left) + node + Inorder(Tree[node].Right)
}
func Postorder(node string) string {
    if node == "." {
        return ""
    }
    return Postorder(Tree[node].Left) + Postorder(Tree[node].Right) + node
}
func main() {
    sc.Split(bufio.ScanWords)
    defer wt.Flush()

    N := nextInt()
    for i := 0; i < N; i++ {
        sc.Scan()
        root := sc.Text()
        sc.Scan()
        left := sc.Text()
        sc.Scan()
        right := sc.Text()
        Tree[root] = Node{Left: left, Right: right}
    }
    wt.WriteString(Preorder("A"))
    wt.WriteByte('\n')
    wt.WriteString(Inorder("A"))
    wt.WriteByte('\n')
    wt.WriteString(Postorder("A"))
}
