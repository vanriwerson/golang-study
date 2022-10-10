package enderecos

import "strings"

// TipoDeEndereco verifica o logradouro e o retorna caso seja um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavra := strings.Split(endereco, " ")[0]
	primeiraPalavra = strings.ToLower(primeiraPalavra)

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			return strings.Title(primeiraPalavra)
		}
	}

	return "Tipo inválido"
}
