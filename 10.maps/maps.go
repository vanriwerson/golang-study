package main

import "fmt"

func main() {
	usuario := map[string]string{ // declara-se o tipo da chave e do valor (devem ser do mesmo tipo)
		"nome":      "Pedro", // os valores de um map não são acessíveis por dot notation
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) // acesso ao valor de um map

	usuario2 := map[string]map[string]string{ // maps aninhados
		"nomeCompleto": {
			"primeiroNome": "Getúlio",
			"sobrenome":    "Vargas",
		},
		"curso": {
			"nome":       "Engenharia",
			"modalidade": "Presencial",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nomeCompleto"]["sobrenome"])

	delete(usuario2, "nomeCompleto")
	fmt.Println(usuario2) // a chave nomeCompleto foi deletada do map
}
