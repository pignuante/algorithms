package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

type Node struct {
    num          int
    before, next *Node
}

type Deck struct {
    head, tail *Node
    size       int
}

func (deck *Deck) PushFront(num int) {
    node := Node{num: num, before: nil, next: nil}
    if deck.size == 0 {
        deck.head = &node
        deck.head.next = &node
        deck.tail = &node
        deck.tail.before = &node
    } else {
        node.next = deck.head
        deck.head.before = &node
        deck.head = &node
    }
    deck.size++
}
func (deck *Deck) PushBack(num int) {
    node := Node{num: num, before: nil, next: nil}
    if deck.size == 0 {
        deck.head = &node
        deck.head.next = &node
        deck.tail = &node
        deck.tail.before = &node
    } else {
        node.before = deck.tail
        deck.tail.next = &node
        deck.tail = &node
    }
    deck.size++
}
func (deck *Deck) PopFront() (num int) {
    if deck.size  == 0 {
        num = -1
    }  else if deck.size == 1{
        num = deck.head.num
        deck.head = nil
        deck.tail = nil
        deck.size--
    } else {
        num = deck.head.num
        deck.head = deck.head.next
        deck.head.before = nil
        deck.size--
    }
    return
}
func (deck *Deck) PopBack() (num int) {
    if deck.size == 0 {
        num = -1
    } else if deck.size == 1{
        num = deck.tail.num
        deck.tail = nil
        deck.head = nil
        deck.size--
    } else {
        num = deck.tail.num
        deck.tail = deck.tail.before
        deck.tail.next = nil
        deck.size--
    }
    return
}

func (deck Deck) Size() int {
    return deck.size
}

func (deck Deck) Empty() (isEmtpy int) {
    if deck.size == 0 {
        isEmtpy = 1
    }
    return
}

func (deck Deck) Front() (num int) {
    if deck.size == 0 {
        num = -1
    } else {
        num = deck.head.num
    }
    return
}
func (deck Deck) Back() (num int) {
    if deck.size == 0 {
        num = -1
    } else {
        num = deck.tail.num
    }
    return
}

func main() {
    wt := bufio.NewWriter(os.Stdout)
    sc := bufio.NewScanner(os.Stdin)
    defer wt.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())
    var deck Deck
    for ; N > 0; N-- {
        sc.Scan()
        com := strings.Fields(sc.Text())
        switch com[0] {
        case "push_front":
            num, _ := strconv.Atoi(com[1])
            deck.PushFront(num)
        case "push_back":
            num, _ := strconv.Atoi(com[1])
            deck.PushBack(num)
        case "pop_front":
            num := deck.PopFront()
            wt.WriteString(strconv.Itoa(num) + "\n")
        case "pop_back":
            num := deck.PopBack()
            wt.WriteString(strconv.Itoa(num) + "\n")
        case "size":
            wt.WriteString(strconv.Itoa(deck.Size()) + "\n")
        case "empty":
            wt.WriteString(strconv.Itoa(deck.Empty()) + "\n")
        case "front":
            num := deck.Front()
            wt.WriteString(strconv.Itoa(num) + "\n")
        case "back":
            num := deck.Back()
            wt.WriteString(strconv.Itoa(num) + "\n")
        }
    }
}
