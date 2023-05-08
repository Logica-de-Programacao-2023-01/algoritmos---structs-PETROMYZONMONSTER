package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	var X float64
	fmt.Println("Digite o valor do raio:")
	fmt.Scan(&X)
	Val := Circulo{X}
	Area := CalcArea(Val)
	fmt.Println("Sua área é:", Area)
}

func CalcArea(Val Circulo) float64 {
	return math.Pi * Val.raio * Val.raio
}
