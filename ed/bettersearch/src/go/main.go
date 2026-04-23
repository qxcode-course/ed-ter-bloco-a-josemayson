package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	esquerda := 0
	direita := len(slice) - 1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if value == slice[meio] {
			return true, meio
		}

		if value > slice[meio] {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}
	return false, esquerda
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
