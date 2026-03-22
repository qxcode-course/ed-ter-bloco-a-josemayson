package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var pessoas []int

	for i := 0; i < n; i++ {
		var aux int
		fmt.Scan(&aux)
		pessoas = append(pessoas, aux)
	}

	var qtd int
	fmt.Scan(&qtd)
	test := make(map[int]bool)

	for i := 0; i < qtd; i++ {
		var aux int
		fmt.Scan(&aux)
		test[aux] = true
	}

	var novoVetor []int
	for _, p := range pessoas {
		if !test[p] {
			novoVetor = append(novoVetor, p)
		}
	}

	for i := range novoVetor {
		fmt.Printf("%d ", novoVetor[i])
	}
	fmt.Println()
}
