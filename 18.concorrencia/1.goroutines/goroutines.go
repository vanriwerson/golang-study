package main

import (
	"fmt"
	"time"
)

func main() {
	// escrever("Sou um loop infinito") // executando desse modo, a segunda chamada da função escrever nunca será feita

	// usando goroutine, as execuções se tornarão concorrentes (o tempo de execução é distribuido entre elas)
	go escrever("Ainda sou um loop infinito, mas estou sendo executado em concorrência") // a palavra reservada go faz com que o programa não espere a execução dessa linha para prosseguir
	escrever("Nunca serei impresso... a menos que você comente a linha 9")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
