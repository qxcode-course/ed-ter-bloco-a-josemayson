package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (set *Set) String() string {
	return "[" + Join(set.data[:set.size], ", ") + "]"
}

func (set *Set) reserve(capacity int) {
	if set.capacity < capacity {
		set.capacity = capacity
	}
}

func (set *Set) insert(index, value int) error {
	if index > set.size {
		return errors.New("index out of range")
	}

	if set.size == set.capacity {
		aux := set.capacity * 2
		if aux == 0 {
			aux = 1
		}
		set.reserve(aux)
	}

	set.size++
	set.data = append(set.data, 0)
	copy(set.data[index+1:], set.data[index:])
	set.data[index] = value
	return nil
}

func (set *Set) erase(value int) bool {
	aux := set.binarySearch(value)
	if aux != -1 {
		set.data = append(set.data[:aux], set.data[aux+1:]...)
		set.size--
		return true
	}
	return false
}

func (set *Set) Contains(value int) bool {
	aux := set.binarySearch(value)
	if aux != -1 {
		return true
	} else {
		return false
	}
}

func (set *Set) binarySearch(value int) int {
	esquerda := 0
	direita := set.size - 1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if value == set.data[meio] {
			return meio
		}

		if value > set.data[meio] {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}
	return -1
}

func (set *Set) Insert(value int) {
	if set.Contains(value) {
		return
	}

	posicao := 0

	for i := range set.data {
		if value < set.data[i] {
			break
		}
		posicao++
	}
	set.insert(posicao, value)
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
