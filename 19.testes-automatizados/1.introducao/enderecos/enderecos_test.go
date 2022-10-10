package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Boa prática: Test+NomeDaFuncTestada
func TestTipoDeEndereco(t *testing.T) {
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
