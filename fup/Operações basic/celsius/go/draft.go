package main
import "fmt"

func main() {
    var celsius float64
    fmt.Scan(&celsius)

    fahrenheit := 1.8 * celsius + 32

    fmt.Printf("%f\n", fahrenheit)
}
