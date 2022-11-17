package main

import (
	"errors"
	"fmt"
)

type person struct {
	name   string
	age    int
	height float32
}

func createPerson(name string, age int, height float32) (error, *person) {
	if len(name) <= 1 || age < 0 || height < 0 {
		return errors.New("Dados invalidos"), nil
	}
	return nil, &person{name, age, height}
}

func main() {

	p1 := person{"Felipe", 20, 1.83}
	fmt.Println("Nome de p1: ", p1.name)

	err, p2 := createPerson("Samuel", -2, 1.73)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p2.name)

}
