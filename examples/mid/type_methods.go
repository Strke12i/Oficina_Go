package main

import "fmt"

type coordenadas struct {
	x int
	y int
}

func soma_coord(a coordenadas, b coordenadas) coordenadas {
	return coordenadas{a.x + b.x, a.y + b.y}
}

func (a coordenadas) soma(b coordenadas) coordenadas {
	return coordenadas{a.x + b.x, a.y + b.y}
}

func main() {
	a := coordenadas{10, 2}
	b := coordenadas{2, 1}

	fmt.Println("soma coord: ", soma_coord(a, b))
	fmt.Println("Metodo soma de a: ", a.soma(b))
}
