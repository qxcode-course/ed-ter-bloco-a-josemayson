package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)

	var d, aux int = 0, math.MaxInt

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
		if a[i] > 9 && b[i] > 9 {
			if aux > int(math.Abs(float64(a[i]-b[i]))) {
				aux = int(math.Abs(float64(a[i] - b[i])))
			}
		} else {
			d++
		}
	}

	if d == n {
		fmt.Println("sem ganhador")
	} else {
		for i := 0; i < n; i++ {
			if a[i] > 9 && b[i] > 9 {
				if int(math.Abs(float64(a[i]-b[i]))) == aux {
					fmt.Println(i)
					break
				}
			}
		}
	}
}
