package main

import "fmt"

type Pessoa struct {
	Name  string
	Email string
	Age   int16
}

type Professor struct {
	Pessoa
	Diploma string
}

type Teste interface {
	int | string | float32
}

func (p Pessoa) mostrar() {
	fmt.Println("Dados da pessoa", p)
}

const STATE = "DONE"

func main() {

	lista := []string{"a", "g", "c"}

	fmt.Println("hello world")

	fmt.Println(lista)

	professor := Professor{
		Pessoa: Pessoa{
			Name:  "leandro",
			Email: "leandro@gmail.cm",
			Age:   15,
		},
		Diploma: "ads",
	}

	professor.mostrar()
}

func add(a int, b int) int {
	return a + b
}
