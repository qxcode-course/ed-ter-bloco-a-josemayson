package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node[T comparable] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
	root  *Node[T]
}

type LList[T comparable] struct {
	root *Node[T]
	size int
}

func (list *Node[T]) Next() *Node[T] {
	if list.next == list.root {
		return list.root.next
	}
	return list.next
}

func (list *Node[T]) Prev() *Node[T] {
	if list.prev == list.root {
		return list.root.prev
	}
	return list.prev
}

func NewLList[T comparable]() *LList[T] {
	root := &Node[T]{}
	root.prev = root
	root.next = root
	root.root = root
	return &LList[T]{root: root, size: 0}
}

func (list *LList[T]) String() string {
	saida := "["

	for it := list.root.next; it != list.root; it = it.next {
		saida += fmt.Sprint(it.Value)
		if it.next != list.root {
			saida += ", "
		}
	}
	return saida + "]"
}

func (list *LList[T]) Insert(node *Node[T], value T) {
	a := node.prev
	c := &Node[T]{
		Value: value,
		next:  node,
		prev:  a,
		root:  list.root,
	}
	a.next = c
	node.prev = c
	list.size++
}

func (list *LList[T]) PushBack(value T) {
	list.Insert(list.root, value)
}

func (list *LList[T]) PushFront(value T) {
	list.Insert(list.root.next, value)
}

func (list LList[T]) Size() int {
	return list.size
}

func (list *LList[T]) Clear() {
	list.root.prev = list.root
	list.root.next = list.root
}

func (list *LList[T]) Search(value T) *Node[T] {
	for it := list.root.next; it != list.root; it = it.next {
		if it.Value == value {
			return it
		}
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList[int]()

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
		case "clear":
			ll.Clear()
		case "forward":
			search, _ := strconv.Atoi(args[1])
			steps, _ := strconv.Atoi(args[2])
			node := ll.Search(search)
			if node == nil {
				fmt.Println("fail: valor não encontrado")
				continue
			}
			collect := []string{}
			for range steps {
				collect = append(collect, fmt.Sprintf("%v", node.Value))
				node = node.Next()
			}
			fmt.Printf("[ %s ]\n", strings.Join(collect, " "))
		case "backward":
			search, _ := strconv.Atoi(args[1])
			steps, _ := strconv.Atoi(args[2])
			node := ll.Search(search)
			if node == nil {
				fmt.Println("fail: valor não encontrado")
				continue
			}
			collect := []string{}
			for range steps {
				collect = append(collect, fmt.Sprintf("%v", node.Value))
				node = node.Prev()
			}
			fmt.Printf("[ %s ]\n", strings.Join(collect, " "))
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
