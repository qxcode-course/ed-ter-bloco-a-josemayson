package main

import "fmt"

func main() {
	var qtd int
	fmt.Scan(&qtd)

	animais := make([]int, qtd)
	solteiros := make(map[int]int)

	for i := range qtd {
		fmt.Scan(&animais[i])
	}

	n_pares := 0
	for _, animal := range animais {
		if solteiros[-animal] > 0 {
			solteiros[-animal]--
			n_pares++
		} else {
			solteiros[animal]++
		}
	}

	fmt.Println(n_pares)
}
