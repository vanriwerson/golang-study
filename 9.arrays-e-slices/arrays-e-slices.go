package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int   // Declaramos o tamanho e o tipo de dados do array
	array1[0] = 121     // É possível popular uma posição específica através do índice
	fmt.Println(array1) // Todos os elementos do array devem ter o mesmo tipo

	array2 := [4]string{"Aqui", "temos", "quatro", "posições"}
	fmt.Println(array2)

	array3 := [...]int{1, 2} // As reticências definem o tamanho do array de acordo com o número de valores passados pro array
	fmt.Print("array3", array3)

	slice := []int{10, 11, 12, 13, 14} // Possui tamanho dinâmico
	fmt.Println(slice)

	fmt.Println("Tipo do slice", reflect.TypeOf(slice)) // Slices são diferentes de arrays
	fmt.Println("Tipo do array3", reflect.TypeOf(array3))

	slice = append(slice, 15)
	fmt.Println(slice)

	// slices são pedaços de um array e trabalha como ponteiro
	slice2 := array2[1:3] // o range escolhido tem início inclusivo e final exclusivo
	fmt.Println(slice2)
}
