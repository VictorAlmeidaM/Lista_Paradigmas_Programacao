package main

import "fmt"

// Estrutura Carro2
type Carro struct {
	marca      string
	modelo     string
	ano        int
	velocidade int
}

// Função construtora para criar um novo carro
func NovoCarro(marca, modelo string, ano int) Carro {
	return Carro{
		marca:      marca,
		modelo:     modelo,
		ano:        ano,
		velocidade: 0, // Velocidade inicial
	}
}

// Método para exibir os detalhes do carro
func (c Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.marca, c.modelo, c.ano)
}

// Método para acelerar
func (c *Carro) Acelerar(incremento int) {
	c.velocidade += incremento
	fmt.Printf("%s %s acelerou. Velocidade atual: %d km/h\n", c.marca, c.modelo, c.velocidade)
}

// Método para frear
func (c *Carro) Frear(decremento int) {
	c.velocidade -= decremento
	if c.velocidade < 0 {
		c.velocidade = 0 // Garantindo que a velocidade não seja negativa
	}
	fmt.Printf("%s %s freou. Velocidade atual: %d km/h\n", c.marca, c.modelo, c.velocidade)
}

// Método para exibir a velocidade atual
func (c Carro) ExibirVelocidade() {
	fmt.Printf("Velocidade atual do %s %s: %d km/h\n", c.marca, c.modelo, c.velocidade)
}

// Função principal para testar a estrutura
func main() {
	// Instanciando três objetos da estrutura Carro
	carro1 := NovoCarro("Toyota", "Yaris", 2020)
	carro2 := NovoCarro("Honda", "Accord", 2019)
	carro3 := NovoCarro("Ford", "Edge", 2018)

	// Exibindo detalhes e controlando a velocidade dos carros
	carro1.ExibirDetalhes()
	carro1.Acelerar(30)
	carro1.Frear(10)
	carro1.ExibirVelocidade()

	carro2.ExibirDetalhes()
	carro2.Acelerar(50)
	carro2.Frear(20)
	carro2.ExibirVelocidade()

	carro3.ExibirDetalhes()
	carro3.Acelerar(40)
	carro3.Frear(50) // Tentativa de frear mais do que a velocidade
	carro3.ExibirVelocidade()
}
