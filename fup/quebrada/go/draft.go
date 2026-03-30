package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Scan(&n1, &n2)

	fmt.Printf("%.0f\n", n1/n2)
	fmt.Printf("%d\n", int(int(n1)%int(n2)))
	fmt.Printf("%.2f\n", n1/n2)
}
