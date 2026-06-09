package main

import "fmt"

func busca(n int, atual int, divisor int) int {
	if divisor*divisor > atual {
		if n == 1 {
			return atual
		}
		return busca(n-1, atual+1, 2)
	}

	if atual%divisor == 0 {
		return busca(n, atual+1, 2)
	}

	return busca(n, atual, divisor+1)
}

func enesimo(n int) int {
	return busca(n, 2, 2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(enesimo(n))
}
