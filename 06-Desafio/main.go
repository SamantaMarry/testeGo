package main

import "fmt"

func main() {

	marvel := make(map[string]int)
	marvel["Homem de Ferro"] = 10
	marvel["Viúva Negra"] = 1000000
	marvel["Homem-Aranha"] = 12
	marvel["Thor"] = 7
	marvel["Capitã Marvel"] = 3

	dc := make(map[string]int)
	dc["Mulher Maravilha"] = 10
	dc["Batman"] = 5
	dc["Super-Homem"] = 9
	dc["Mulher-Gato"] = 7
	dc["Hera"] = 3

	heroi := "Homem de Ferro"
	determinaValor(marvel[heroi])

	heroi = "Batman"
	determinaValor(dc[heroi])

}

func determinaValor(nivel int) {
	// É fraco 7 < >= 7 é Forte
	if nivel < 7 {
		fmt.Println("Fraco")
	} else {
		fmt.Println("Muito Forte")
	}

}