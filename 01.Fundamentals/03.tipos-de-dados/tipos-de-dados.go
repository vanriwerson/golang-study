package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÚMEROS INTEIROS:
	var inteiro int8 = -100 // Não definir o número de bits faz com que o go utilize a arquitetura da máquina em que está
	fmt.Println(inteiro)

	var inteiro2 uint = 10000 // Não admite sinal
	fmt.Println(inteiro2)

	// alias int32 = rune
	var inteiro3 rune = 12345
	fmt.Println(inteiro3)

	// alias uint8 = byte
	var inteiro4 rune = 123
	fmt.Println(inteiro4)

	// NÚMEROS REAIS:
	var real float32 = 123.45
	fmt.Println(real)

	var real2 float64 = 123000.45
	fmt.Println(real2)

	real3 := 123456.78
	fmt.Println(real3)

	// CADEIA DE CARACTERES
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char) //retorna o número da tabela ASCII com o tipo int

	// VALOR ZERO: é atribuído a uma variável não inicializada
	var texto string
	var num int

	fmt.Println(texto)
	fmt.Println(num)

	// BOOLEAN
	var boolean bool // o valor zero é false
	fmt.Println(boolean)

	// ERROR
	var erro error // o valor zero é <nil>
	fmt.Println(erro)

	var declaredError error = errors.New("Mensagem do tipo erro")
	fmt.Println(declaredError)
}
