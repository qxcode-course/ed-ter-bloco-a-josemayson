package main

import "fmt"

func imprimir(temp []int) {
	for i := range temp {
		fmt.Printf("%d", temp[i])
		if i < len(temp) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
func main() {
	var qtd int
	fmt.Scan(&qtd)

	temp := make(map[int]bool)
	fila := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		fmt.Scan(&fila[i])
	}

	var novafila []int

	var qtdsair int
	fmt.Scan(&qtdsair)

	sair := make([]int, qtdsair)

	for i := range sair {
		fmt.Scan(&sair[i])
	}

	for _, p := range sair {
		temp[p] = true
	}

	for _, p := range fila {
		if !temp[p] {
			novafila = append(novafila, p)
			temp[p] = true
		}
	}

	imprimir(novafila)

}
