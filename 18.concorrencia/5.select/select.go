package main

import (
	"fmt"
	"time"
)

// O select 'é um switch' voltado para a performance do uso de concorrência
func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}() // Chamada da função anônima

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select { // faz com que o canal que estiver pronto para enviar a mensagem execute essa ação de modo independente
		case mensagemCanal1 := <-canal1: // sem o select, esse canal teria sua performance prejudicada pelo canal2
			fmt.Println(mensagemCanal1) // teria que esperar os 2 segundos, mesmo podendo enviar uma mensagem a cada meio segundo

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
