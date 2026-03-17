package main

import "fmt"

func main() {
	var qtdTotal, qtdPossui int
	fmt.Scan(&qtdTotal, &qtdPossui)

	figurinhas := make([]int, qtdPossui)

	for i := 0; i < qtdPossui; i++ {
		fmt.Scan(&figurinhas[i])
	}

	album := make([]int, qtdTotal+1)
	var figurinhasRepetidas []int

	for i := 0; i < qtdPossui; i++ {
		aux := figurinhas[i]

		if album[aux] > 0 {
			figurinhasRepetidas = append(figurinhasRepetidas, aux)
		}
		album[aux]++
	}

	tamanhoRepetidas := len(figurinhasRepetidas)

	if tamanhoRepetidas == 0 {
		fmt.Println("N")
	} else {
		for i := 0; i < tamanhoRepetidas; i++ {
			fmt.Printf("%d", figurinhasRepetidas[i])

			if i < tamanhoRepetidas-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	var faltando []int

	for i := 1; i <= qtdTotal; i++ {
		if album[i] == 0 {
			faltando = append(faltando, i)
		}
	}

	tamanhoFaltando := len(faltando)

	if tamanhoFaltando == 0 {
		fmt.Println("N")
	} else {
		for i := 0; i < tamanhoFaltando; i++ {
			fmt.Printf("%d", faltando[i])

			if i < tamanhoFaltando-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
