package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c + 1},
		{p.l, p.c - 1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	q := NewQueue[Pos]()
	q.Enqueue(startPos)

	visited := make(map[Pos]bool)
	visited[startPos] = true

	caminho := make(map[Pos]Pos)

	for !q.IsEmpty() {
		curr, _ := q.Dequeue()

		if curr == endPos {
			break
		}

		for _, vizinho := range curr.getNeig() {
			if inside(grid, vizinho) && grid[vizinho.l][vizinho.c] != '#' && !visited[vizinho] {
				visited[vizinho] = true
				q.Enqueue(vizinho)
				caminho[vizinho] = curr
			}
		}
	}
	voltar(grid, startPos, endPos, caminho)
}

func voltar(grid [][]rune, startPos Pos, endPos Pos, caminho map[Pos]Pos) {
	if _, exist := caminho[endPos]; !exist && startPos != endPos {
		return
	}
	curr := endPos

	for curr != startPos {
		grid[curr.l][curr.c] = '.'
		curr = caminho[curr]
	}
	grid[startPos.l][startPos.c] = '.'

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
