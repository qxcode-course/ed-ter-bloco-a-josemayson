package main

import (
	"fmt"
	"math"
)

type jogada struct {
	pa, pb int
}

func abs(a int, b int) int {
	if a-b >= 0 {
		return a - b
	} else {
		return (a - b) * (-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	jogadas := make([]jogada, n)

	var aux, maior int = 0, math.MaxInt

	for i := range n {
		fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)
		if jogadas[i].pa > 9 && jogadas[i].pb > 9 {
			diferenca := abs(jogadas[i].pa, jogadas[i].pb)
			if maior > diferenca {
				maior = diferenca
			}
		} else {
			aux++
		}
	}

	if aux == n {
		fmt.Println("sem ganhador")
	} else {
		for i := range n {
			if jogadas[i].pa > 9 && jogadas[i].pb > 9 {
				auxiliar := abs(jogadas[i].pa, jogadas[i].pb)
				if maior == auxiliar {
					fmt.Println(i)
					break
				}
			}
		}
	}

}
