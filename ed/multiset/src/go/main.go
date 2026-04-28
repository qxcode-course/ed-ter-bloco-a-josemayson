package main

import (
	"bufio"
	"errors"
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
		data:     make([]int, 0, capacity),
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

func (set *MultiSet) search(value int) (bool, int) {
	esquerda := 0
	direita := set.size - 1
	resultado := -1
	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if value == set.data[meio] {
			esquerda = meio + 1
			resultado = meio
		} else if value > set.data[meio] {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}
	if resultado != -1 {
		return true, resultado
	} else {
		return false, esquerda
	}
}

func (set *MultiSet) Contains(value int) bool {
	aux, _ := set.search(value)
	return aux
}

func (set *MultiSet) insert(index, value int) error {
	if index > set.size {
		return errors.New("index out of range")
	}

	if set.capacity == set.size {
		set.expand()
	}
	set.size++
	set.data = append(set.data, 0)
	copy(set.data[index+1:], set.data[index:])
	set.data[index] = value
	return nil
}

func (set *MultiSet) Insert(value int) {
	esquerda := 0
	direita := set.size - 1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if value >= set.data[meio] {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}

	set.insert(esquerda, value)
}

func (set *MultiSet) String() string {
	return "[" + Join(set.data, ", ") + "]"
}

func (set *MultiSet) Count(value int) int {
	temp := 0
	for i := range set.data {
		if set.data[i] == value {
			temp++
		}
	}
	return temp
}

func (set *MultiSet) Clear() {
	set.data = set.data[:0]
	set.size = 0
}

func (set *MultiSet) unique() int {
	if set.size == 0 {
		return 0
	}
	temp := 1

	for i := 0; i < set.size-1; i++ {
		if set.data[i] != set.data[i+1] {
			temp++
		}
	}
	return temp
}

func (set *MultiSet) erase(index int) error {
	if set.size < index {
		return errors.New("value not found")
	}
	set.data = append(set.data[:index], set.data[index+1:]...)
	set.size--
	return nil
}

func (set *MultiSet) Erase(value int) error {
	indice := -1

	for i := range set.data {
		if value == set.data[i] {
			indice = i
			break
		}
	}
	if indice != -1 {
		set.erase(indice)
		return nil
	}
	return errors.New("value not found")
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
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
