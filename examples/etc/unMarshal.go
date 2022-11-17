package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}
type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := Record{
		"Felipe",
		"Colpo",
		[]Telephone{Telephone{true, "128883393989"}},
	}
	rec, err := json.Marshal(&myRecord)
	if err != nil {
		return
	}
	fmt.Println(string(rec))

	var unRec Record
	err = json.Unmarshal(rec, &unRec)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(unRec)
}
