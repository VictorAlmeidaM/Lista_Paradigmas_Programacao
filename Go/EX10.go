package main

import "fmt"

// Definindo a estrutura Calculadora
type Calculadora struct{}

// Função somar para dois números
func (c Calculadora) SomarDois(num1, num2 int) int {
    return num1 + num2
}

// Função somar para três números
func (c Calculadora) SomarTres(num1, num2, num3 int) int {
    return num1 + num2 + num3
}

func main() {
    // Criando um objeto da Calculadora
    calc := Calculadora{}

    // Soma de dois números
    resultado1 := calc.SomarDois(10, 20)
    fmt.Printf("Soma de 10 + 20: %d\n", resultado1)

    // Soma de três números
    resultado2 := calc.SomarTres(10, 20, 30)
    fmt.Printf("Soma de 10 + 20 + 30: %d\n", resultado2)
}
