package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(board [][]byte, word string) bool {
	linhas := len(board)
	colunas := len(board[0])
	total := linhas * colunas

	var busca func(r, c, i int) bool
	busca = func(r, c, i int) bool {
		if len(word) == i {
			return true
		}

		if c < 0 || r < 0 || r >= linhas || c >= colunas || board[r][c] != word[i] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		aux := busca(r+1, c, i+1) || busca(r-1, c, i+1) || busca(r, c+1, i+1) || busca(r, c-1, i+1)
		board[r][c] = temp
		return aux
	}

	for i := 0; i < total; i++ {
		r := i / colunas
		c := i % colunas

		if busca(r, c, 0) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
