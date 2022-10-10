package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) { // t.Run é análogo ao describe
		ret := Retangulo{10, 12}
		esperado := float64(120)
		recebido := ret.Area()

		if esperado != recebido {
			t.Fatalf("A área recebida %f é diferente da área esperada %f.", recebido, esperado) // Fatalf para a execução dos testes
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		esperado := float64(math.Pi * 100)
		recebido := circ.Area()

		if esperado != recebido {
			t.Fatalf("A área recebida %f é diferente da área esperada %f.", recebido, esperado)
		}
	})
}
