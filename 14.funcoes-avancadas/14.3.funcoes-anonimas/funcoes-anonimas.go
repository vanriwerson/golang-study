package main

import "fmt"

func main() {
	func() {
		fmt.Println("Minha função anônima")
	}() // Os parênteses fazem o papel de invocar a função

	saudacao := func(texto string) string { // sintaxe análoga às arrow functions do javascript
		return fmt.Sprintf("Saudações, %s!", texto)
	}("Invocador") // Esses parênteses também recebem os parâmetros

	fmt.Println(saudacao)
}
