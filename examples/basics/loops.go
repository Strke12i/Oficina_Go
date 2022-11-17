package main

import "fmt"

func main() {

	fmt.Println("Loop For")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Exemplo de while true")
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

}
