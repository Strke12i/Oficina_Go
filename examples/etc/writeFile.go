package main

import (
	"fmt"
	"io"
	"os"
)

func writeFile(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := writeFile("teste.txt", "Texto escrito"); err != nil {
		fmt.Println("Error ao escrever")
		return
	}
	fmt.Println("Escrito com sucesso")

}
