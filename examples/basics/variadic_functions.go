package main

import "fmt"

func contador(nums ...int) int {
	cont := 0
	for _, _ = range nums {
		cont++
	}
	return cont
}

func main() {
	c := contador(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Total: ", c)
}
