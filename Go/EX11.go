package main

import (
	"fmt"
)

// Definição da interface Funcionario
type Funcionario interface {
	CalcularSalario() float64
	GetNome() string
}

// Estrutura FuncionarioHorista
type FuncionarioHorista struct {
	nome             string
	horasTrabalhadas int
	valorHora        float64
}

// Construtor para FuncionarioHorista
func NewFuncionarioHorista(nome string, horasTrabalhadas int, valorHora float64) *FuncionarioHorista {
	return &FuncionarioHorista{nome: nome, horasTrabalhadas: horasTrabalhadas, valorHora: valorHora}
}

// Implementação do método CalcularSalario
func (f *FuncionarioHorista) CalcularSalario() float64 {
	return float64(f.horasTrabalhadas) * f.valorHora
}

// Implementação do método GetNome
func (f *FuncionarioHorista) GetNome() string {
	return f.nome
}

// Estrutura FuncionarioAssalariado
type FuncionarioAssalariado struct {
	nome          string
	salarioMensal float64
}

// Construtor para FuncionarioAssalariado
func NewFuncionarioAssalariado(nome string, salarioMensal float64) *FuncionarioAssalariado {
	return &FuncionarioAssalariado{nome: nome, salarioMensal: salarioMensal}
}

// Implementação do método CalcularSalario
func (f *FuncionarioAssalariado) CalcularSalario() float64 {
	return f.salarioMensal
}

// Implementação do método GetNome
func (f *FuncionarioAssalariado) GetNome() string {
	return f.nome
}

// Função principal para teste
func main() {
	// Criando um funcionário horista
	funcionarioHorista := NewFuncionarioHorista("Carlos", 203, 15)
	fmt.Printf("Salário de %s: R$%.2f\n", funcionarioHorista.GetNome(), funcionarioHorista.CalcularSalario())

	// Criando um funcionário assalariado
	funcionarioAssalariado := NewFuncionarioAssalariado("Melissa", 3000)
	fmt.Printf("Salário de %s: R$%.2f\n", funcionarioAssalariado.GetNome(), funcionarioAssalariado.CalcularSalario())
}
