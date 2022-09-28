package main

import "fmt"

func main() {
	fmt.Println("======= OPERADORES ARITMÉTICOS =======")
	soma := 1 + 2
	fmt.Println("1 + 2 =", soma)

	subtracao := 5 - 2
	fmt.Println("5 - 2 =", subtracao)

	divisao := 9 / 3
	fmt.Println("9 / 3 =", divisao)

	mutiplicacao := 3 * 1
	fmt.Println("3 * 1 =", mutiplicacao)

	modulo := 10 % 7
	fmt.Println("10 % 7 =", modulo)

	// em Go, para realizar qualquer tipo de operação entre dados, eles DEVEM ser do MESMO TIPO
	var num1 int16 = 10
	var num2 int16 = 25     // mudar o tipo declarado para int32 (exemplo) geraria um erro
	novaSoma := num1 + num2 // se os tipos fossem diferentes a soma não seria realizada

	fmt.Println("Operação com variáveis: 10 + 25 =", novaSoma)

	fmt.Println("====== OPERADORES DE ATRIBUIÇÃO ======")
	var texto string = "Atribuição com tipo explícito"
	texto2 := "Atribuição com inferência de tipo"

	fmt.Println(texto)
	fmt.Println(texto2)

	fmt.Println("======= OPERADORES RELACIONAIS =======")
	fmt.Println(num1 < num2) // true
	fmt.Println(num1 > num2) // false
	fmt.Println(7 == 8)      // false
	fmt.Println(int16(soma) == int16(modulo))
	fmt.Println(7 <= 7)       // true
	fmt.Println(num1 != num2) // true

	fmt.Println("========= OPERADORES LÓGICOS =========")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // And - retorna false
	fmt.Println(verdadeiro || falso) // Or - retorna true
	fmt.Println(!verdadeiro)         // Not - retorna false

	fmt.Println("========= OPERADORES UNÁRIOS =========")
	num3 := 10
	num3++
	fmt.Println("num3++ =>", num3) // retorna 11
	num3 += 9
	fmt.Println("num3 += 9 =>", num3) // retorna 20
	num3 -= 10
	fmt.Println("num3 -= 10 =>", num3) // retorna 10
	num3--
	fmt.Println("num3-- =>", num3) // retorna 9
	num3 /= 3
	fmt.Println("num3 /= 3 =>", num3) // retorna 3
	num3 *= 10
	fmt.Println("num3 *= 10 =>", num3) // retorna 30
	num3 %= 7
	fmt.Println("num3 %= 7 =>", num3) // retorna 2

	// Go não possui operador ternário.
	var ternarioRaiz string
	if num3 > 5 {
		ternarioRaiz = "Maior que 5"
	} else {
		ternarioRaiz = "Menor que 5"
	}

	fmt.Println("Estrutura if/else num3 = 2 =>", ternarioRaiz)
}
