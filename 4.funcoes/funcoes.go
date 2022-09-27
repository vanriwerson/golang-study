package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Funções em Go podem ter mais de um retorno
func calcular(n1, n2 int8) (int8, int8) {
	subtracao := n1 - n2
	divisao := n1 / n2

	return subtracao, divisao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("Função f " + txt)
		return txt
	}

	execute := f("com parâmetro")
	fmt.Println(execute)

	subtracao, divisao := calcular(30, 10) // é necessário haver tantas variáveis quanto os retornos da função
	// a ordem dos retornos é sempre a que foi gerada na implementação
	fmt.Println(subtracao, divisao)

	// é possivel utilizar o underScore '_' para ignorar retornos
	_, apenasDivisao := calcular(30, 2)
	fmt.Println(apenasDivisao)
}
