package main

import "fmt"

func changePointer(n *int) {
	*n = 10
}

func main() {
	number := 1
	fmt.Println("Endereço de number: ", &number)
	fmt.Println("Valor de number: ", number)

	changePointer(&number)
	fmt.Println(number)
}
