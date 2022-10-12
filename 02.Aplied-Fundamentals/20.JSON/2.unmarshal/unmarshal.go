package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`
	var c1 cachorro

	// unmarshal recebe um slice de bytes que representa o JSON que será convertido e
	// o endereço de memória (ponteiro) da variável que receberá o resultado da conversão
	if erro := json.Unmarshal([]byte(cachorroJSON), &c1); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c1) // retorna em formato de struct

	cachorro2JSON := `{"nome":"Tobby","raca":"Poodle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2JSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2) // retorna em formato de map
}
