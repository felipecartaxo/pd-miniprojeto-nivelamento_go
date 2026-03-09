package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func showMenu() {
	fmt.Println("\n=== Gerenciador de Números ===")
	fmt.Println("1) Adicionar número")
	fmt.Println("0) Sair")
	fmt.Print("Escolha uma opção: ")
}

func addNum(sliceNums []int) []int {
	fmt.Print("Digite um número inteiro para adicionar: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler input:", err)
		return sliceNums
	}

	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Valor inválida. Digite um número inteiro.")
		return sliceNums
	}

	sliceNums = append(sliceNums, num)
	fmt.Printf("Número %d adicionado com sucesso!\n", num)

	return sliceNums
}

// -----
func main() {
	sliceNums := []int{}

	for {
		showMenu()

		optionStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler opção:", err)
			continue
		}
		optionStr = strings.TrimSpace(optionStr)

		switch optionStr {
		case "1":
			sliceNums = addNum(sliceNums)
		case "0":
			fmt.Println("Encerrando aplicação...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}