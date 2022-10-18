package modelos

import "time"

type Usuario struct {
	ID       uint64    `json:"id,omitempty"` // omit empty faz com que o campo não seja passado para o JSON caso esteja vazio
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}
