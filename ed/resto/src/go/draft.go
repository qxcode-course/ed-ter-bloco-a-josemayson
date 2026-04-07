package main

import "fmt"

func imprimir(inteiro int) {
	if inteiro < 1 {
		return
	}
	aux := inteiro / 2
	temp := inteiro % 2
	imprimir(inteiro / 2)
	fmt.Println(aux, temp)
}
func main() {
	var inteiro int
	fmt.Scan(&inteiro)

	imprimir(inteiro)
}
