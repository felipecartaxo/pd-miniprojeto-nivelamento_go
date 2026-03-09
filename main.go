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
	fmt.Println("2) Listar números")
	fmt.Println("3) Remover por índice")
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

func listNums(sliceNums []int) {
	if len(sliceNums) == 0 {
		fmt.Println("A lista está vazia.")
		return
	}

	fmt.Println("Números armazenados:")
	for i, num := range sliceNums {
		fmt.Printf("Índice %d: %d\n", i, num)
	}
}

func removeByIndex(sliceNums []int) []int {
	if len(sliceNums) == 0 {
		fmt.Println("A lista está vazia. Nada para remover.")
		return sliceNums
	}

	fmt.Print("Digite o índice que deseja remover: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler input:", err)
		return sliceNums
	}

	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Índice inválido. Digite um número inteiro.")
		return sliceNums
	}

	if index < 0 || index >= len(sliceNums) {
		fmt.Println("Índice fora do intervalo da lista.")
		return sliceNums
	}

	removedValue := sliceNums[index]

	sliceNums = append(sliceNums[:index], sliceNums[index+1:]...)

	fmt.Printf("Número %d removido com sucesso do índice %d!\n", removedValue, index)
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
		case "2":
			listNums(sliceNums)
		case "3":
			sliceNums = removeByIndex(sliceNums)
		case "0":
			fmt.Println("Encerrando aplicação...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}