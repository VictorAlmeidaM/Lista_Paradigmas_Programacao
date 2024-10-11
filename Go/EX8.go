package main

import (
	"fmt"
)

// Estrutura Empregado
type Empregado struct {
	nome   string
	cargo  string
	salario float64
}

// Construtor para Empregado
func NovoEmpregado(nome, cargo string, salario float64) Empregado {
	return Empregado{
		nome:   nome,
		cargo:  cargo,
		salario: salario,
	}
}

// Método para exibir informações do empregado
func (e Empregado) ExibirInformacoes() {
	fmt.Printf("Nome: %s, Cargo: %s, Salário: R$%.2f\n", e.nome, e.cargo, e.salario)
}

// Estrutura Empresa
type Empresa struct {
	nomeEmpresa string
	empregados  []Empregado // Lista para armazenar os empregados da empresa
}

// Construtor para Empresa
func NovaEmpresa(nomeEmpresa string) Empresa {
	return Empresa{
		nomeEmpresa: nomeEmpresa,
		empregados:  []Empregado{},
	}
}

// Método para adicionar um empregado
func (e *Empresa) AdicionarEmpregado(empregado Empregado) {
	e.empregados = append(e.empregados, empregado)
}

// Método para listar os empregados da empresa
func (e Empresa) ListarEmpregados() {
	fmt.Printf("Lista de empregados da empresa %s:\n", e.nomeEmpresa)
	for _, empregado := range e.empregados {
		empregado.ExibirInformacoes()
	}
}

// Método para calcular a folha de pagamento
func (e Empresa) CalcularFolhaPagamento() {
	totalSalarios := 0.0
	for _, empregado := range e.empregados {
		totalSalarios += empregado.salario
	}
	fmt.Printf("Folha de pagamento total: R$%.2f\n", totalSalarios)
}

// Função principal
func main() {
	// Criando empregados
	emp1 := NovoEmpregado("Joaquim", "Desenvolvedor", 5000)
	emp2 := NovoEmpregado("Marisa", "Gerente", 7500)
	emp3 := NovoEmpregado("Pedro", "Designer", 4000)

	// Criando uma empresa e adicionando empregados
	empresa := NovaEmpresa("Tech Solutions")
	empresa.AdicionarEmpregado(emp1)
	empresa.AdicionarEmpregado(emp2)
	empresa.AdicionarEmpregado(emp3)

	// Exibindo a lista de empregados e calculando a folha de pagamento
	empresa.ListarEmpregados()
	empresa.CalcularFolhaPagamento()
}
