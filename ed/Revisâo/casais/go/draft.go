package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	animal := make([]int, n)
	descasados := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&animal[i])
		if animal[i] < 0 {
			descasados[i] = animal[i]
		}
	}

	var t int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if animal[i] > 0 {
				if animal[i] == (descasados[j] * (-1)) {
					descasados[j] = 0
					t++
				}
			}
		}
	}
	fmt.Println(t)
}
