package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Número de argumentos invalidos")
		return
	}

	var soma float64
	soma = 0.0

	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		number, _ := strconv.ParseFloat(arguments[i], 64)
		soma += number
	}

	fmt.Println("O resultado da soma é: ", soma)

}
