package main

import "fmt"

func main() {
	var Q int
	fmt.Scan(&Q)

	var D string
	fmt.Scan(&D)

	x := make([]int, Q)
	y := make([]int, Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	for i := 0; i < Q-1; i++ {
		x[Q-i-1] = x[Q-i-2]
		y[Q-i-1] = y[Q-i-2]
	}

	switch D {
	case "R":
		x[0] += 1
	case "D":
		y[0] += 1
	case "U":
		y[0] -= 1
	case "L":
		x[0] -= 1
	}

	for i := 0; i < Q; i++ {
		fmt.Printf("%d %d\n", x[i], y[i])
	}

}
