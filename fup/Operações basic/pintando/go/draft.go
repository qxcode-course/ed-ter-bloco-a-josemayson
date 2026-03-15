package main
import (
    "fmt"
    "math"
)

func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    perimetro := (a + b + c) / 2.0
    area := math.Sqrt(perimetro * ( perimetro - a) * (perimetro - b) * (perimetro - c))

    fmt.Printf("%.2f\n", area)

}
