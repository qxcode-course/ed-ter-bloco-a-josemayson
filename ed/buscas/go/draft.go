package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	entrada := make([]string, n)

	for i := range entrada {
		fmt.Scan(&entrada[i])
	}

	var x int
	fmt.Scan(&x)

	consulta := make([]string, x)

	for i := range consulta {
		fmt.Scan(&consulta[i])
	}

	temp := make(map[string]int)

	for _, p := range entrada {
		temp[p]++
	}

	for i := range consulta {
		fmt.Printf("%d", temp[consulta[i]])
		if i < len(consulta)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
