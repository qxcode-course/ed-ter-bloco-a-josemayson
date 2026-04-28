package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	esquerda := 0
	direita := len(slice) - 1
	resultado := -1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if value == slice[meio] {
			esquerda = meio + 1
			resultado = meio
		}
		if value < slice[meio] {
			direita = meio - 1
		} else {
			esquerda = meio + 1
		}
	}

	if resultado != -1 {
		return resultado
	}
	return esquerda
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
