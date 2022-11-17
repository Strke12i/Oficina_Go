package main

import "fmt"

func main() {
	valor_inteiro := 10
	valor_float := 19.98
	valor_complexo := complex(12, 1)

	fmt.Printf("Inteiro %T\n", valor_inteiro)
	fmt.Printf("Float %T\n", valor_float)
	fmt.Printf("Numero Complexp %T\n", valor_complexo)
}
