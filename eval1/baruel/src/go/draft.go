package main

import "fmt"

func imprimir(temp []int) {
	for i := range temp {
		fmt.Printf("%d", temp[i])
		if i < len(temp)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	var e int
	fmt.Scan(&e)
	album := make([]int, e)
	var repetidas []int
	var faltando []int
	aux := make(map[int]bool)

	for i := 0; i < e; i++ {
		fmt.Scan(&album[i])
	}

	for _, p := range album {
		if aux[p] {
			repetidas = append(repetidas, p)
		}
		aux[p] = true
	}

	for i := 1; i <= n; i++ {
		if !aux[i] {
			faltando = append(faltando, i)
		}
	}

	if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		imprimir(repetidas)
	}

	if len(faltando) == 0 {
		fmt.Println("N")
	} else {
		imprimir(faltando)
	}

}
