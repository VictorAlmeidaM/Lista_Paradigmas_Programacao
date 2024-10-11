package main

import (
	"fmt"
)

// Definindo a struct Matematica
type Matematica struct{}

// Método que calcula o fatorial de um número
func (m Matematica) Fatorial(n int) int {
	if n < 0 {
		panic("O fatorial não está definido para números negativos.")
	} else if n == 0 || n == 1 {
		return 1
	} else {
		resultado := 1
		for i := 2; i <= n; i++ {
			resultado *= i
		}
		return resultado
	}
}

// Função principal para teste
func main() {
	m := Matematica{}
	numero := 4
	resultado := m.Fatorial(numero)
	fmt.Printf("O fatorial de %d é: %d\n", numero, resultado)
}
