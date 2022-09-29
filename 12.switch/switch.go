package main

import "fmt"

func diaDaSemana(num int) string {
	switch num { // em Go não temos a cláusula break
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "O valor informado não corresponde a um dia da semana"
	}
}

func eFimDeSemana(num int) string {
	switch {
	case num == 1 || num == 7:
		return "É fim de semana!!!"
	case num > 1 && num < 4:
		return "Não. Acabamos de começar a semana."
	case num >= 4 && num < 7:
		return "Calma, você consegue chegar lá."
	default:
		return "O valor informado não corresponde a um dia da semana"
	}
}

func comFallThrough(num int) string {
	var diaDaSemana string

	switch {
	case num == 1:
		diaDaSemana = "Domingo"
		fallthrough // joga para a próxima condição sem avaliar
	case num == 2:
		diaDaSemana = "Segunda-feira"
	case num == 3:
		diaDaSemana = "Terça-feira"
	case num == 4:
		diaDaSemana = "Quarta-feira"
	case num == 5:
		diaDaSemana = "Quinta-feira"
	case num == 6:
		diaDaSemana = "Sexta-feira"
	case num == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "O valor informado não corresponde a um dia da semana"
	}

	return diaDaSemana
}

func main() {
	fmt.Println(diaDaSemana(5))  // Quinta-feira
	fmt.Println(diaDaSemana(51)) // default

	fmt.Println(eFimDeSemana(1))
	fmt.Println(eFimDeSemana(3))
	fmt.Println(eFimDeSemana(4))
	fmt.Println(eFimDeSemana(7))
	fmt.Println(eFimDeSemana(71))

	fmt.Println(comFallThrough(1))
}
