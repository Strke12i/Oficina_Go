package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Quantidade de argumentos inválidos")
		return
	}

	palavra := os.Args[1]

	fmt.Println("Caracter  | Valor ")
	for i := 0; i < len(palavra); i++ {
		fmt.Printf("%c       | %d \n", palavra[i], palavra[i])
	}
	// Todo : Fazer caso para caracteres fora do padrão ascii
}
