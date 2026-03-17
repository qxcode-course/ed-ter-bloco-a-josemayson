package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	semi := (a + b + c) / 2.0
	area := math.Sqrt(semi * (semi - a) * (semi - b) * (semi - c))

	fmt.Printf("%.2f\n", area)
}
