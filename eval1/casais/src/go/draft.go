package main

import "fmt"

func main() {
	var qtd int
	fmt.Scan(&qtd)

	animal := make([]int, qtd)

	for i := range animal {
		fmt.Scan(&animal[i])
	}

	casais := make(map[int]bool)
	n := 0

	for _, p := range animal {
		if casais[-p] {
			n++
			casais[-p] = false
		} else {
			casais[p] = true
		}
	}

	fmt.Println(n)
}
