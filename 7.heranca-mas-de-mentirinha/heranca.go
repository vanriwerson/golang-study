package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct { // todo estudante é uma pessoa
	// pessoa pessoa // criaria uma chave pessoa dentro da struct estudante
	pessoa    // estudante herda as características da struct pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Eric", "Johnson", 46, 178}
	fmt.Println(p1)

	fmt.Println("O tipo estudante herda as características do tipo pessoa:")
	e1 := estudante{p1, "Música", "Berkley"}
	fmt.Println(e1)
}
