package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (vet *Vector) Status() string {
	return fmt.Sprintf("size:%v capacity:%v", vet.size, vet.capacity)
}

func (vet *Vector) String() string {
	return "[" + Join(vet.data[0:vet.size], ", ") + "]"
}

func (vet *Vector) Reserve(newCapacity int) {
	vet.capacity = newCapacity
}

func (vet *Vector) PushBack(value int) {
	if vet.size == vet.capacity {
		if vet.capacity == 0 {
			vet.capacity++
		} else {
			vet.capacity *= 2
		}
		newData := make([]int, vet.capacity)
		copy(newData, vet.data)
		vet.data = newData
	}
	vet.data[vet.size] = value
	vet.size++
}

func (vet *Vector) Capacity() int {
	return vet.capacity
}

func (vet *Vector) Size() int {
	return vet.size
}

func (vet *Vector) Clear() {
	vet.size = 0
}

func (vet *Vector) IndexOf(value int) int {
	for i := range vet.data {
		if value == vet.data[i] {
			return i
		}
	}
	return -1
}

func (vet *Vector) Contains(value int) bool {
	for i := range vet.data {
		if value == vet.data[i] {
			return true
		}
	}
	return false
}

func (vet *Vector) PopBack() error {
	if vet.size == 0 {
		return errors.New("vector is empty")
	}
	vet.size--
	return nil
}

func (vet *Vector) Get(index int) int {
	return vet.data[index]
}

func (vet *Vector) At(index int) (int, error) {
	if index >= vet.size {
		return 0, errors.New("index out of range")
	}
	return vet.data[index], nil
}

func (vet *Vector) Set(index, value int) error {
	if index >= vet.size || index < 0 {
		return errors.New("index out of range")
	}
	vet.data[index] = value
	return nil
}

func (vet *Vector) Insert(index, value int) error {
	if vet.capacity <= vet.size {
		vet.capacity *= 2
	}
	vet.data = append(vet.data, 0)
	vet.size++
	copy(vet.data[index+1:], vet.data[index:])
	vet.data[index] = value
	return nil
}

func (vet *Vector) Erase(index int) error {
	if index > vet.capacity {
		return errors.New("index out of range")
	}
	vet.data = append(vet.data[:index], vet.data[index+1:]...)
	vet.size--
	return nil
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

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			// start, _ := strconv.Atoi(parts[1])
			// end, _ := strconv.Atoi(parts[2])
			// slice := v.Slice(start, end)
			// fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
