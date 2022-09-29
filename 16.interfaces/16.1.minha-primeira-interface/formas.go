package main

import (
	"fmt"
	"math"
)

type forma interface { // interfaces entregam mais flexibilidade quanto aos tipos
	area() float64 // forma exige um método area retornando float64
}

func escreverArea(f forma) { // esse método pertence à interface forma
	fmt.Printf("A área da forma é %0.2f.\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 { // criação do método area para o struct retangulo
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 { // criação do método area para o struct circulo
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r) // inicialização do método area através da interface forma

	c := circulo{2}
	escreverArea(c)
}
