package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero float64
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func EndPes(pessoa Pessoa) Pessoa {
	return Pessoa{
		Nome:  "Lucas",
		Idade: 20,
		Endereco: Endereco{
			Rua:    "Rua 10",
			Numero: 16,
			Estado: "SP",
		},
	}
}

func main() {
	P := EndPes(Pessoa{})
	fmt.Printf("%s,Número: %g,Estado: %s", P.Endereco.Rua, P.Endereco.Numero, P.Endereco.Estado)
}
