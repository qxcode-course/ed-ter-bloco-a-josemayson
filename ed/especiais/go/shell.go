package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	qtd := make(map[int]int)

	var contagem []Pair

	for _, p := range vet {
		if p < 0 {
			p *= -1
		}
		qtd[p]++
	}
	for i, p := range qtd {
		contagem = append(contagem, Pair{One: i, Two: p})
	}

	sort.Slice(contagem, func(i, j int) bool {
		return contagem[i].One < contagem[j].One
	})
	return contagem
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair{}
	}
	qtd := 0
	temp := vet[0]
	for i := range vet {
		if vet[i] < 0 {
			vet[i] *= -1
		}
	}
	var time []Pair
	for i := range vet {
		if vet[i] == temp {
			qtd++
		} else {
			time = append(time, Pair{One: temp, Two: qtd})
			temp = vet[i]
			qtd = 1
		}
	}
	time = append(time, Pair{One: temp, Two: qtd})
	return time
}

func mnext(vet []int) []int {
	lista := make([]int, len(vet))

	for i := range vet {
		if vet[i] > 0 {
			if i < len(vet)-1 {
				if vet[i+1] < 0 {
					lista[i] = 1
				}
			}
		}
		if i > 0 {
			if vet[i-1] < 0 && vet[i] > 0 {
				lista[i] = 1
			}
		}
	}
	return lista
}

func alone(vet []int) []int {
	lista := make([]int, len(vet))

	for i := range vet {
		temp := false
		if vet[i] > 0 {
			if i < len(vet)-1 {
				if vet[i+1] < 0 {
					temp = true
				}
			}
		}
		if i > 0 {
			if vet[i-1] < 0 && vet[i] > 0 {
				temp = true
			}
		}
		if vet[i] > 0 {
			if temp == false {
				lista[i] = 1
			}
		}
	}
	return lista
}

func couple(vet []int) int {
	casais := make(map[int]bool)
	qtd := 0

	for _, p := range vet {
		if casais[-p] {
			casais[-p] = false
			qtd++
		} else {
			casais[p] = true
		}
	}
	return qtd
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	j := 0
	for i := 0; i < len(vet); i++ {
		if vet[i] == seq[j] {
			j++
			if j == len(seq) {
				return i - len(seq) + 1
			}
		} else {
			if j > 0 {
				i -= j
				j = 0
			}
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	var novalista []int
	indice := make(map[int]bool)

	for _, p := range posList {
		indice[p] = true
	}

	for i := range vet {
		if !indice[i] {
			novalista = append(novalista, vet[i])
		}
	}
	return novalista
}

func clear(vet []int, value int) []int {
	lista := make(map[int]bool)
	var novaLista []int
	lista[value] = true
	for _, p := range vet {
		if !lista[p] {
			novaLista = append(novaLista, p)
		}
	}
	return novaLista
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
