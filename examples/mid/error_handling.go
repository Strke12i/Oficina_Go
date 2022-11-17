package main

import (
	"errors"
	"fmt"
	"os"
)

func verificaLenArgument() error {
	if len(os.Args) < 2 {
		return errors.New("número de argumentos insuficiente")
	}
	return nil
}

func main() {
	err := verificaLenArgument()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ultimo argumento:", os.Args[len(os.Args)-1])
}
