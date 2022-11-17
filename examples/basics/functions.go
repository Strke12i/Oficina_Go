package main

import "fmt"

func retornaAnteriores(num int) (int, int) {
	return num - 1, num + 1
}

func retornaPointer(x int) *int {
	y := x + 1
	return &y
}

func main() {
	ant, pos := retornaAnteriores(10)
	fmt.Println("", ant, "- 10 -", pos)

	func_dobro := func(x int) int {
		return 2 * x
	}
	fmt.Println("O dobro de 5 é:", func_dobro(5))

	pointer := retornaPointer(10)
	fmt.Println("Valor de pointer: ", *pointer)
	fmt.Println("Endereço de pointer: ", pointer)

}
