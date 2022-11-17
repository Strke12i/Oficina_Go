package main

import (
	"fmt"
	"sort"
)

func main() {
	Slice := make([]int, 5) //inicializa o slice com 20 valores de 0
	//Adiciona 4 valores no slice
	for i := 5; i > 0; i-- {
		Slice = append(Slice, i)
	}

	for _, i := range Slice {
		fmt.Println(i)
	}

	fmt.Println("Tamanho do slice ", len(Slice))

	reSlice := Slice[5:9]
	fmt.Println("reSlice ", reSlice)

	sort.Slice(reSlice, func(i, j int) bool {
		return reSlice[i] < reSlice[j]
	})

	fmt.Println("reSlice sorted ", reSlice)
}
