package main

import "fmt"

// Estrutura Motor
type Motor struct {
	potencia         int    // em cavalos de potência (cv)
	tipoCombustivel  string // gasolina, etanol, etc.
}

// Método para exibir informações do motor
func (m *Motor) exibirInformacoes() {
	fmt.Printf("Motor: %d cv, Tipo de Combustível: %s\n", m.potencia, m.tipoCombustivel)
}

// Estrutura Carro
type Carro struct {
	marca   string
	modelo  string
	ano     int
	motor   Motor // Atributo que é um objeto da estrutura Motor
}

// Método para exibir detalhes do carro
func (c *Carro) exibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.marca, c.modelo, c.ano)
	c.motor.exibirInformacoes() // Exibe informações do motor associado
}

// Função principal
func main() {
	// Criando um objeto Motor
	motor1 := Motor{potencia: 150, tipoCombustivel: "Gasolina"}

	// Criando um objeto Carro associado ao Motor
	carro1 := Carro{marca: "Toyota", modelo: "Yaris", ano: 2023, motor: motor1}

	// Exibindo detalhes do carro e do motor
	carro1.exibirDetalhes()
}
