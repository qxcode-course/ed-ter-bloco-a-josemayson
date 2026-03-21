package main

import "fmt"

func main() {
	var qtdtotal, qtdpossui int
	fmt.Scan(&qtdtotal, &qtdpossui)

	aux := make(map[int]bool)
	var repetidas []int
	var faltando []int

	for i := 0; i < qtdpossui; i++ {
		var temp int
		fmt.Scan(&temp)
		if aux[temp] {
			repetidas = append(repetidas, temp)
		}
		aux[temp] = true
	}

	for i := 1; i <= qtdtotal; i++ {
		if !aux[i] {
			faltando = append(faltando, i)
		}
	}

	if len(repetidas) > 0 {
		for i, num := range repetidas {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	} else {
		fmt.Println("N")
	}

	if len(faltando) > 0 {
		for i, num := range faltando {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	} else {
		fmt.Println("N")
	}
}
