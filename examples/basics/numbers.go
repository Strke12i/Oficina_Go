package main

import "fmt"

func main() {

	var signed_int8 int8
	signed_int8 = -127 // de -128 até 127
	// signed_int8 = 128 ?

	var unsigned_int8 uint8
	unsigned_int8 = 10 // 0 até 127
	// unsigned_int8 = -9 ?

	valor_inteiro := 10
	valor_float := 19.98
	valor_complexo := complex(12, 1)

	fmt.Printf("SignedInt8: %T\n", signed_int8)
	fmt.Printf("UnsignedInt: %T\n", unsigned_int8)
	fmt.Printf("Inteiro %T\n", valor_inteiro)
	fmt.Printf("Float %T\n", valor_float)
	fmt.Printf("Numero Complexo %T\n", valor_complexo)
}
