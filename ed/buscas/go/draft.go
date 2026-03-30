package main

import "fmt"

func main() {
	var qtdentrada int
	fmt.Scan(&qtdentrada)

	entrada := make([]string, qtdentrada)
	for i := range qtdentrada {
		fmt.Scan(&entrada[i])
	}

	var qtdconsulta int
	fmt.Scan(&qtdconsulta)
	consulta := make([]string, qtdconsulta)
	for i := range qtdconsulta {
		fmt.Scan(&consulta[i])
	}

}
