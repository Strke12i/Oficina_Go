package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file *os.File
	file = os.Stdin
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Out>", scanner.Text())
	}
}
