package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func main() {
	p := Pessoa{
		Nome:  "Lucas",
		Idade: 20,
		Endereco: Endereco{
			Rua:    "Rua 10",
			Numero: 16,
			Estado: "SP",
		},
	}
	fmt.Println(p)
}
