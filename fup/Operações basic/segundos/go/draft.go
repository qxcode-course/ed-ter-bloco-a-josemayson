package main
import "fmt"

func main() {
    var tempo int
    fmt.Scan(&tempo)
    
    hora := tempo / 3600
    aux := tempo % 3600
    minutos := aux / 60
    segundos := aux % 60

    fmt.Printf("%d:%d:%d\n", hora, minutos, segundos)
}
