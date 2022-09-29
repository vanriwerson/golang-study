package main

import "fmt"

// Go não é uma linguagem orientada a objetos

type usuario struct { // Structs são tipos personalizados
	nome     string //Respeitam o princípio do 'valor zero'
	idade    uint8
	endereco endereco // é possível inserir structs dentro de structs
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	var u usuario // criando struct com o tipo explícito
	u.nome = "Renato"
	u.idade = 18
	fmt.Println(u) // retornará o valor zero para a chave endereco

	end := endereco{"Destinos", 10000}
	u2 := usuario{"Humberto", 51, end} // com inferência de tipo
	fmt.Println(u2)

	u3 := usuario{nome: "Eric"} // com apenas um valor informado
	fmt.Println(u3)             // retornará o valor zero para as chaves idade e endereco
}
