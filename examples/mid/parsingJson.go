package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("teste.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var parsedData map[string]interface{}
	json.Unmarshal([]byte(file), &parsedData)

	for key, value := range parsedData {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}
