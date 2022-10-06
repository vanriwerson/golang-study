package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // canais podem enviar e receber dados
	go escrever("Olá, mundo!", canal)

	fmt.Println("Início do programa.") // aparecerá antes da go routine

	// for {
	// 	mensagem, aberto := <-canal // (await) recebendo valor pelo canal
	// 	if !aberto {
	// 		break // encerra o loop caso o canal esteja fechado
	// 	}

	// 	fmt.Println(mensagem) // depois dessa linha o programa é finalizado sem executar todas as repetições
	// }

	// refactoring infinite loop
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa.")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // (async) enviando o valor pelo canal
		time.Sleep(time.Second)
	}

	close(canal) // evita o erro fatal deadlock (canal esperando um valor que nunca chegará)
}
