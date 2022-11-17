package main

import "fmt"

func AlteraMapa(mapa *map[string]map[string]int) {
	for _, m := range *mapa {
		if m["X"] == 12 {
			m["Y"] = 54
		}
	}
}

func main() {
	mapa := map[string]int{
		"X": 10,
		"Y": 20,
		"Z": 30,
	}

	for key, value := range mapa {
		fmt.Printf("%s:[%d]\n", key, value)
	}

	fmt.Println("Mapa com o valor da Chave X: ", mapa["X"])

	if _, err := mapa["T"]; err {
		fmt.Println("Mapa na posição T existe")
	} else {
		fmt.Println("Mapa na posição T não existe")
	}

	mapa2 := map[string]map[string]int{
		"uid1": {
			"X": -109,
			"Y": 98,
		},
		"uid": {
			"X": 12,
			"Y": -54,
		},
	}

	fmt.Println(mapa2)
	AlteraMapa(&mapa2)
	fmt.Println(mapa2)

}
