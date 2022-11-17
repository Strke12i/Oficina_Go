package main

import "fmt"

func main() {
	Array1D := [3]int{1, 2, 3}
	Array2D := [3][2]int{{0, 1}, {2, 3}, {4, 5}}

	fmt.Println("Len Array1D ", len(Array1D))
	fmt.Println("Len Array2D", len(Array2D))

	//Print com index
	fmt.Println("Array com uma dimensão")
	for index, valor := range Array1D {
		fmt.Println(index, valor)
	}

	fmt.Println("Array com duas dimensões")
	for _, valor := range Array2D {
		for _, i := range valor {
			fmt.Println(i)
		}
	}

}
