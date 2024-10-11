package main

import (
	"fmt"
)

// Interface Animal
type Animal interface {
	Som() string
}

// Struct Cachorro
type Cachorro struct{}

// Implementação do método Som para Cachorro
func (c Cachorro) Som() string {
	return "Au Au"
}

// Struct Gato
type Gato struct{}

// Implementação do método Som para Gato
func (g Gato) Som() string {
	return "Miau"
}

// Função principal
func main() {
	// Criando instâncias das structs
	var cachorro Animal = Cachorro{}
	var gato Animal = Gato{}

	// Exibindo os sons dos animais
	fmt.Println("Cachorro:", cachorro.Som())
	fmt.Println("Gato:", gato.Som())
}
