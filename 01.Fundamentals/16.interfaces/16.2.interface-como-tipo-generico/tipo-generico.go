package main

import (
	"fmt"
)

func generica(interf interface{}) { // como nada está sendo passado para a interface, qualquer tipo a satisfaz
	fmt.Println(interf)
}

func main() {
	generica("Podemos utilizar uma string")
	generica(12)
	generica(true)
	generica(3.1416)

	/*
		O fmt.Println recebe interfaces genéricas como parâmetro.
		Por isso podemos imprimir qualquer quantidade de elementos de qual tipo.
	*/

	// Um exemplo do que NÃO deve ser feito:
	mapa := map[interface{}]interface{}{
		1235678:                         "string",
		false:                           float32(12),
		"NaoSeiOQueEstaAcontecendoAqui": true,
	}

	fmt.Println(mapa)
}
