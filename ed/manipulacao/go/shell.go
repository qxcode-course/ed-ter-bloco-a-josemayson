package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
		if vet[i] > -10 && vet[i] < 0 {
			mulheres = append(mulheres, vet[i])
		}
	}
	return mulheres
}

func sortVet(vet []int) []int {
	sort.Slice(vet, func(i, j int) bool {
		return vet[i] < vet[j]
	})
	return vet
}

func sortStress(vet []int) []int {
	sort.Slice(vet, func(i, j int) bool {
		return math.Abs(float64(vet[i])) < math.Abs(float64(vet[j]))
	})
	return vet
}

func reverse(vet []int) []int {
	novaLista := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		novaLista[i] = vet[len(vet)-i-1]
	}
	return novaLista
}

func unique(vet []int) []int {
	var novaLista []int
	repetidos := make(map[int]bool)

	for _, p := range vet {
		if !repetidos[p] {
			novaLista = append(novaLista, p)
		}
		repetidos[p] = true
	}
	return novaLista
}

func repeated(vet []int) []int {
	var novaLista []int
	repetidos := make(map[int]bool)

	for _, p := range vet {
		if repetidos[p] {
			novaLista = append(novaLista, p)
		}
		repetidos[p] = true
	}
	return novaLista
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
