package main

import "fmt"

type gomos struct {
	x, y int
}

func main() {
	var Q int
	fmt.Scan(&Q)

	var D string
	fmt.Scan(&D)

	gomo := make([]gomos, Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&gomo[i].x, &gomo[i].y)
	}

	for i := 0; i < Q-1; i++ {
		gomo[Q-i-1].x = gomo[Q-i-2].x
		gomo[Q-i-1].y = gomo[Q-i-2].y
	}

	switch D {
	case "R":
		gomo[0].x += 1
	case "D":
		gomo[0].y += 1
	case "U":
		gomo[0].y -= 1
	case "L":
		gomo[0].x -= 1
	}

	for i := 0; i < Q; i++ {
		fmt.Printf("%d %d\n", gomo[i].x, gomo[i].y)
	}

}
