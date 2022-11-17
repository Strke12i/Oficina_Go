package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println()
	text := "Oficina de Go sainf"

	upper := s.ToUpper(text)
	lower := s.ToLower(text)

	split := s.Split(text, " ")

	fmt.Println("MAISCULO: ", upper)
	fmt.Println("minisculo: ", lower)
	fmt.Println("Contagem de letras a: ", s.Count(text, "a"))
	fmt.Println("Palavras em text: ")
	for _, i := range split {
		fmt.Printf("[%s] ", i)
	}
	fmt.Println()

}
