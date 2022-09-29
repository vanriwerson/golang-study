package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// Métodos estão obrigatoriamente relacionados a alguma estrutura
func (u usuario) salvar() { // u = acesso às chaves do struct, usuario = estrutura ao qual o método está relacionado
	fmt.Printf("Salvando os dados de %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversário() { // *usuario aponta para o endereço de memória e altera o dado diretamente
	u.idade += 1
}

func main() {
	usuario1 := usuario{"Gustavo", 29}
	usuario1.salvar() // o método pode ser acessado por dot notation
	fmt.Println("maior de idade:", usuario1.maiorDeIdade())

	usuario2 := usuario{"Guanabara", 17}
	usuario2.salvar()
	fmt.Println("maior de idade:", usuario2.maiorDeIdade())

	fmt.Println("Executando método fazerAniversário()")
	usuario2.fazerAniversário() // esse método acessa o endereço de memória onde está armazenado usuario2
	fmt.Println("maior de idade:", usuario2.maiorDeIdade())
}
