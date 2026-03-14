package main
import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    var d, e ,f float64
    fmt.Scan(&d, &e, &f)
    var g float64
    fmt.Scan(&g)

    p := float64(a) * d + float64(b) * e + f * float64(c)

    fmt.Printf("%.2f\n", g - p)
}
