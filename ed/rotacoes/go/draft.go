package main

import "fmt"

func main() {
	var t, r int
	fmt.Scan(&t, &r)

	var lista []int

	for i := 1; i <= t; i++ {
		lista = append(lista, i)
	}

	for i := 0; i < r; i++ {
		aux := lista[len(lista)-1]
		copy(lista[1:], lista[:len(lista)-1])
		lista[0] = aux
	}
	fmt.Print("[ ")
	for i := range lista {
		fmt.Printf("%d ", lista[i])
	}
	fmt.Print("]\n")

}
