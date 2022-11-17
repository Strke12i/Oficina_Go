package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile(text []string) error {
	file, err := os.Create("escrita.txt")

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	for _, s := range text {
		fmt.Println(s)
		writer.WriteString(s + "\n")
	}

	writer.Flush()
	file.Close()
	return nil
}

func main() {

	file, err := ioutil.ReadFile("nomes.json")
	if err != nil {
		fmt.Println("Não foi possivel ler o arquivo!")
		return
	}

	var nomes map[string]interface{}
	err = json.Unmarshal([]byte(file), &nomes)
	if err != nil {
		fmt.Println("Impossivel fazer o parse!")
		return
	}

	var text []string

	for _, t := range nomes {
		field := t.(map[string]interface{})
		nome := field["nome"]
		idade := field["idade"].(float64)
		curso := field["curso"]
		text = append(text, fmt.Sprintf("Me chamo %s, tenho %d anos e curso %s", nome, int(idade), curso))
	}

	err = writeFile(text)
	if err != nil {
		fmt.Println("Impossivel escrever no arquivo")
	}

}
