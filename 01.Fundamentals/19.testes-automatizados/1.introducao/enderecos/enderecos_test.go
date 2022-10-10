package enderecos_test // go aceita que os testes e funcs principais pertençam ao mesmo pacote

import (
	. "introducao-testes/enderecos" // import alias
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Boa prática: Test+NomeDaFuncTestada
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // Indica que esse teste pode ser rodado em paralelo com outros testes que possuem essa chamada

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paluista", "Avenida"},
		{"Estrada dos Coiotes", "Estrada"},
		{"Rodovia Anhanguera", "Rodovia"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Tipo inválido"},
		{"Praça da Amizade", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if cenario.retornoEsperado != retornoRecebido {
			t.Errorf("O tipo recebido %s é diferente do esperado %s.",
				retornoRecebido, cenario.retornoEsperado)
		}
	}
}

// go test ./... faz com que o go procure e execute todos os testes do diretório atual e subdiretórios
// go test -v (verbose) traz mais detalhes obre o teste em execução
// go test --cover traz a porcentagem de cobertura de testes
// go test --coverprofile nomeDoArquivo.txt gera um relatório com a cobertura de testes
// go tool cover --func=nomeDoArquivo.txt lista as funções testadas
// go tool cover --html=nomeDoArquivo.txt destaca as linhas não cobertas pelos testes
