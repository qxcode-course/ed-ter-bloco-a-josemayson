package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var homens []int
	for i := range vet {
		if vet[i] > 0 {
			homens = append(homens, vet[i])
		}
	}
	return homens
}

func getCalmWomen(vet []int) []int {
	var mulheres []int
	for i := range vet {
		if vet[i] < 0 && vet[i] > -10 {
			mulheres = append(mulheres, vet[i])
		}
	}
	return mulheres
}

func sortVet(vet []int) []int {
	for i := range vet {
		for j := i + 1; j < len(vet); j++ {
			if vet[i] > vet[j] {
				vet[i], vet[j] = vet[j], vet[i]
			}
		}
	}
	return vet
}

func sortStress(vet []int) []int {
	for i := range vet {
		for j := i + 1; j < len(vet); j++ {
			if int(math.Abs(float64(vet[i])-0)) > int(math.Abs(float64(vet[j])-0)) {
				vet[i], vet[j] = vet[j], vet[i]
			}
		}
	}
	return vet
}

func reverse(vet []int) []int {
	var lista []int
	for i := range vet {
		aux := vet[len(vet)-1-i]
		lista = append(lista, aux)
	}
	return lista
}

func unique(vet []int) []int {
	lista := make(map[int]bool)
	var aux []int
	for i, p := range vet {
		if !lista[p] {
			aux = append(aux, vet[i])
		}
		lista[p] = true
	}
	return aux
}

func repeated(vet []int) []int {
	lista := make(map[int]bool)
	var aux []int
	for i, p := range vet {
		if lista[p] {
			aux = append(aux, vet[i])
		}
		lista[p] = true
	}
	return aux
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
