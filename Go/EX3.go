package main

import "fmt"

// Estrutura ContaBancaria
type ContaBancaria struct {
	titular string
	saldo   float64 // Atributo protegido por encapsulamento
}

// Função construtora para criar uma nova conta bancária com saldo inicial
func NovaContaBancaria(titular string, saldoInicial float64) ContaBancaria {
	return ContaBancaria{
		titular: titular,
		saldo:   saldoInicial,
	}
}

// Sobrecarga do construtor para saldo inicial opcional
func NovaContaBancariaComSaldoZero(titular string) ContaBancaria {
	return NovaContaBancaria(titular, 0) // Chama o construtor principal com saldo inicial = 0
}

// Método para depositar dinheiro
func (c *ContaBancaria) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
		fmt.Printf("Depósito de R$%.2f realizado com sucesso. Saldo atual: R$%.2f\n", valor, c.saldo)
	} else {
		fmt.Println("Valor de depósito inválido.")
	}
}

// Método para sacar dinheiro
func (c *ContaBancaria) Sacar(valor float64) {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		fmt.Printf("Saque de R$%.2f realizado com sucesso. Saldo atual: R$%.2f\n", valor, c.saldo)
	} else {
		fmt.Println("Saldo insuficiente ou valor de saque inválido.")
	}
}

// Método para exibir o saldo
func (c ContaBancaria) ExibirSaldo() {
	fmt.Printf("Saldo atual de %s: R$%.2f\n", c.titular, c.saldo)
}

// Função principal para testar a estrutura
func main() {
	// Exemplo de uso:
	conta1 := NovaContaBancaria("José", 1000)
	conta1.ExibirSaldo()
	conta1.Depositar(500)
	conta1.Sacar(300)
	conta1.ExibirSaldo()

	conta2 := NovaContaBancariaComSaldoZero("Marina")
	conta2.ExibirSaldo()
	conta2.Sacar(500) // Tentativa de saque maior que o saldo
	conta2.Depositar(100)
	conta2.Sacar(50)
	conta2.ExibirSaldo()
}
