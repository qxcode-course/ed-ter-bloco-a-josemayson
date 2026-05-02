package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

func (no *Node) Prev() *Node {
	if no.prev == no.root {
		return nil
	}
	return no.prev
}

func (no *Node) Next() *Node {
	if no.next == no.root {
		return nil
	}
	return no.next
}

type LList struct {
	head *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.head = &Node{}
	list.head.prev = list.head
	list.head.next = list.head
	list.head.root = list.head
	list.size = 0
	return list
}

func (list *LList) String() string {
	saida := "["
	for it := list.head.next; it != list.head; it = it.next {
		saida += fmt.Sprint(it.Value)
		if it.next != list.head {
			saida += ", "
		}
	}
	return saida + "]"
}

func (list *LList) Insert(b *Node, value int) {
	a := b.prev
	c := &Node{
		Value: value,
		prev:  a,
		next:  b,
		root:  list.head,
	}
	a.next = c
	b.prev = c
	list.size++
}

func (list *LList) PushBack(value int) {
	list.Insert(list.head, value)
}

func (list *LList) PushFront(value int) {
	list.Insert(list.head.next, value)
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
	list.Remove(list.head.prev)
}

func (list *LList) PopFront() {
	list.Remove(list.head.next)
}

func (list *LList) Search(value int) *Node {
	for it := list.head.next; it != list.head; it = it.next {
		if it.Value == value {
			return it
		}
	}
	return nil
}

func (list *LList) Back() *Node {
	if list.size == 0 {
		return nil
	}
	return list.head.prev
}

func (list *LList) Front() *Node {
	if list.size == 0 {
		return nil
	}
	return list.head.next
}

func (list *LList) Clear() {
	list.head.prev = list.head
	list.head.next = list.head
	list.size = 0
}

func (list *LList) Size() int {
	return list.size
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
