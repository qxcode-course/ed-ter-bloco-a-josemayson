package main

import "fmt"

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	var fila []int

	for i := range n {
		fila = append(fila, i+1)
	}
	aux := 0
	for i, p := range fila {
		if p == e {
			aux = i
		}
	}

	for {
		fmt.Printf("[ ")
		for i := range fila {
			if aux == i {
				fmt.Printf("%d> ", fila[i])
			} else {
				fmt.Printf("%d ", fila[i])
			}
		}
		fmt.Printf("]")
		fmt.Println()
		if len(fila) == 1 {
			break
		}
		indice := (aux + 1) % len(fila)
		fila = append(fila[:indice], fila[indice+1:]...)
		aux = indice % len(fila)
	}
}
