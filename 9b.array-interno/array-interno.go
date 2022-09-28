package main

import "fmt"

func main() {
	slice := make([]float32, 10, 11) // slice pega 10 de 11 elementos
	fmt.Println("Slice inicial", slice)
	fmt.Println("Length criado", len(slice))     // Length
	fmt.Println("Capacidade criada", cap(slice)) // Capacidade

	slice = append(slice, 11)
	slice = append(slice, 12)                // Aqui é criado um array interno para não permitir que ultrapassemos a
	fmt.Println("Slice após appends", slice) // capacidade informada inicialmente na função make
	fmt.Println("Novo length", len(slice))
	fmt.Println("Nova capacidade", cap(slice))
}
