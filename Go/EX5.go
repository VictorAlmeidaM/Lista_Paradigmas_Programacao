package main

import "fmt"

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

// Função que recebe uma lista de objetos Animal
func emitirSons(animais []Animal) {
	for _, animal := range animais {
		fmt.Println(animal.Som())
	}
}

// Função principal
func main() {
	// Criando instâncias das structs
	cachorro1 := Cachorro{}
	gato1 := Gato{}
	cachorro2 := Cachorro{}

	// Criando uma lista de animais
	animais := []Animal{cachorro1, gato1, cachorro2}

	// Chamando a função que emite os sons
	emitirSons(animais)
}
