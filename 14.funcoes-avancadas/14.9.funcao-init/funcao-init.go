package main

import "fmt"

var n int

func main() {
	fmt.Println("Executando função main")
	fmt.Println("Esse valor foi setado na função init:", n) // sem a função init, n teria seu valor zero
}

func init() {
	fmt.Println("Executando função init")
	n = 10 // inicializando a variável n
}

// A função init sempre é executada antes da main
// pode haver uma função init por arquivo
// utilizada para fazer setups
