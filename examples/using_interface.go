package main

import (
	"awesomeProject/examples/pack/shape_interface"
	"fmt"
	"math"
)

type quadrado struct {
	lado float32
}
type circulo struct {
	raio float32
}

func (q quadrado) Area() float32 {
	return q.lado * q.lado
}

func (q quadrado) Perimetro() float32 {
	return q.lado * 4
}

func (c circulo) Area() float32 {
	return c.raio * c.raio * math.Pi
}

func (c circulo) Perimetro() float32 {
	return 2 * c.raio * math.Pi
}

func Calculate(x shape_interface.Shape) {
	_, ok := x.(quadrado)
	if ok {
		fmt.Println("É um quadrado")
	} else {
		fmt.Println("É um circulo")
	}
	fmt.Println("Area:", x.Area())
	fmt.Println("Perimetro:", x.Perimetro())
}

func main() {
	q := quadrado{4}
	c := circulo{5}

	fmt.Println("Circulo Area:", c.Area(), "Perimetro:", c.Perimetro())
	fmt.Println("Quadrado Area:", q.Area(), "Perimetro:", q.Perimetro())

}
