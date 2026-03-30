package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Printf("%d\n", a+b)
	fmt.Printf("%d\n", a-b)
	fmt.Printf("%d\n", a*b)
	fmt.Printf("%.2f\n", float64(a)/float64(b))
	fmt.Printf("%d\n", a%b)
}
