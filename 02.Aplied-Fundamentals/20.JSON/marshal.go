package main

import (
	"bytes"
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
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cToJson, erro := json.Marshal(c) // retorna um slice de bytes
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cToJson)
	fmt.Println(bytes.NewBuffer(cToJson)) // 'traduz' os bytes

	c2 := map[string]string{
		"nome": "Tobby",
		"raca": "Poodle",
	}
	fmt.Println(c2)

	c2ToJson, erro := json.Marshal(c2) // retorna um array de bytes
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2ToJson)
	fmt.Println(bytes.NewBuffer(c2ToJson))
}
