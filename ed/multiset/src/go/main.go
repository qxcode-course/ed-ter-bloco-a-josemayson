package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (set *MultiSet) expand() {
	if set.size == set.capacity {
		set.capacity *= 2
	}
	if set.capacity == 0 {
		set.capacity = 1
	}
}

func (set *MultiSet) Insert(value int) {
	if set.size == set.capacity {
		if set.capacity == 0 {
			set.capacity = 1
		} else {
			set.capacity *= 2
		}
	}
	set.data = append(set.data, value)
	set.size++
}

func (set *MultiSet) String() string {
	return "[" + Join(set.data, ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Print(ms)
		case "erase":
			// value, _ := strconv.Atoi(args[1])
		case "contains":
			// value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
