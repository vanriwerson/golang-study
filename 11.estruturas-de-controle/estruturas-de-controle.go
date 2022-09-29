package main

import "fmt"

func main() {
	num := 15

	if num > 15 {
		fmt.Println("Maior do que 15")
	} else if num < 15 {
		fmt.Println("Menor do que 15")
	} else {
		fmt.Println("Exatamente 15")
	}

	num2 := -3
	// IF INIT inicia uma variável no escopo do if
	if init := num2; init > 0 {
		fmt.Println("Maior do que zero")
	} else {
		fmt.Println("Menor do que zero")
	}

	// fmt.Println(init) => gera um erro pois a variável init não existe fora do escopo em que foi criada
}
