package main

import "fmt"

type Triangulo struct {
	base   int
	altura int
}

func main() {
	var n3 Triangulo
	n1 := n3.base
	n2 := n3.altura
	fmt.Println("Digite o valor da base:")
	fmt.Scan(&n1)
	fmt.Println("Digite o valor da altura:")
	fmt.Scan(&n2)
	T := areatri(n1, n2)
	fmt.Println("Sua área é:", T)
}

func areatri(n1 int, n2 int) int {
	return (n1 * n2) / 2
}
