package main

import (
    "bufio"
    "os"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextInt() (r int) {
    sc.Scan()
    for _, c := range sc.Bytes() {
        r *= 10
        r += int(c - '0')
    }
    return
}

var str = [...]string{
    "\"재귀함수가 뭔가요?\"",
    "\"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.",
    "마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.",
    "그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어.\"",
}
var underDash = "____"
var wt *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer wt.Flush()

    N := nextInt()
    wt.WriteString("어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다.")
    wt.WriteByte('\n')
    Recursive(N, "")
}

func Recursive(N int, under string) {
    if N == 0 {
        wt.WriteString(under + "\"재귀함수가 뭔가요?\"")
        wt.WriteByte('\n')
        wt.WriteString(under + "\"재귀함수는 자기 자신을 호출하는 함수라네\"")
        wt.WriteByte('\n')
        wt.WriteString(under + "라고 답변하였지.")
        wt.WriteByte('\n')
        return
    }
    for _, v := range str {
        wt.WriteString(under + v)
        wt.WriteByte('\n')
    }
    Recursive(N-1, under+underDash)
    wt.WriteString(under + "라고 답변하였지.")
    wt.WriteByte('\n')
}
