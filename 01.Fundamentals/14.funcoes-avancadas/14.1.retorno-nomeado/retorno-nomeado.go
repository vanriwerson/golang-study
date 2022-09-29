package main

import "fmt"

func calculos(num1, num2 int) (soma int, subtracao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	return
}

func main() {
	fmt.Println(calculos(12, 5)) // não permite concatenação com strings

	soma, subtracao := calculos(12, 5) // deve-se criar uma variável para cada retorno da função
	fmt.Println("12 + 5 =", soma)
	fmt.Println("12 - 5 =", subtracao)
}
