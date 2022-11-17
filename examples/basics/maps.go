package main

import "fmt"

func main() {
	mapa := map[string]int{
		"X": 10,
		"Y": 20,
		"Z": 30,
	}

	for key, value := range mapa {
		fmt.Printf("%s:[%d]\n", key, value)
	}

	fmt.Println("Mapa com o valor da Chave X: ", mapa["X"])

	if _, err := mapa["T"]; err {
		fmt.Println("Mapa na posição T existe")
	} else {
		fmt.Println("Mapa na posição T não existe")
	}

}
