package main

import "fmt"

// Estrutura Carro
type Carro struct {
	marca  string
	modelo string
	ano    int
}

// Função construtora para criar um novo carro
func NovoCarro(marca, modelo string, ano int) Carro {
	return Carro{
		marca:  marca,
		modelo: modelo,
		ano:    ano,
	}
}

// Método para exibir os detalhes do carro
func (c Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.marca, c.modelo, c.ano)
}

// Função principal
func main() {
	// Instanciando três objetos da estrutura Carro
	carro1 := NovoCarro("Toyota", "Yaris", 2020)
	carro2 := NovoCarro("Honda", "Accord", 2019)
	carro3 := NovoCarro("Ford", "Edge", 2018)

	// Exibindo os detalhes de cada carro
	carro1.ExibirDetalhes()
	carro2.ExibirDetalhes()
	carro3.ExibirDetalhes()
}
