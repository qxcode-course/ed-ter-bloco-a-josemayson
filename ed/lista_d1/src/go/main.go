package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	info int
	next *Node
	prev *Node
}

type LList struct {
	head *Node
	size int
}

func NewLList() *LList {
	head := &Node{}
	head.next = head
	head.prev = head
	return &LList{head: head, size: 0}
}

func (list *LList) Inserir(b *Node, value int) {
	a := b.prev
	c := &Node{
		info: value,
		next: b,
		prev: a,
	}
	a.next = c
	b.prev = c
	list.size++
}

func (list *LList) String() string {
	saida := "["
	for i := list.head.next; i != list.head; i = i.next {
		saida += fmt.Sprint(i.info)
		if i.next != list.head {
			saida += ", "
		}
	}
	return saida + "]"
}

func (list *LList) PushBack(value int) {
	list.Inserir(list.head, value)
}

func (list *LList) PushFront(value int) {
	list.Inserir(list.head.next, value)
}

func (list *LList) Remove(b *Node) {
	a := b.prev
	c := b.next
	a.next = c
	c.prev = a
	b.prev = nil
	b.next = nil
	list.size--
}

func (list *LList) PopBack() {
	if list.head == list.head.next {
		return
	}
	list.Remove(list.head.prev)
}

func (list *LList) PopFront() {
	if list.head == list.head.next {
		return
	}
	list.Remove(list.head.next)
}

func (list *LList) Size() int {
	return list.size
}

func (list *LList) Clear() {
	list.head.next = list.head
	list.head.prev = list.head
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
