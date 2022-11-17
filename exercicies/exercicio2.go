package main

import "fmt"

func soma_valores(arr *[10]int) {
	contador := 0
	for _, i := range arr {
		contador += i
	}

	arr[0] = contador
}

func main() {
	test := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Valor do array antes de chamar a função:", test)

	soma_valores(&test)
	fmt.Println("Valor da soma:", test[0])
	fmt.Println("Valor do array após chamar a função", test)

}
