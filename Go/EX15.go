package main

import (
	"errors"
	"fmt"
)

// SaldoInsuficienteError é uma estrutura de erro personalizada
type SaldoInsuficienteError struct {
	SaldoAtual float64
	ValorSaque float64
}

// Error implementa a interface error para SaldoInsuficienteError
func (e *SaldoInsuficienteError) Error() string {
	return fmt.Sprintf("Saldo insuficiente: R$%.2f para o saque de R$%.2f.", e.SaldoAtual, e.ValorSaque)
}

// ContaBancaria representa uma conta bancária
type ContaBancaria struct {
	Titular string
	Saldo   float64
}

// Depositar adiciona um valor ao saldo da conta
func (c *ContaBancaria) Depositar(valor float64) error {
	if valor <= 0 {
		return errors.New("O valor do depósito deve ser positivo")
	}
	c.Saldo += valor
	fmt.Printf("Depósito de R$%.2f realizado com sucesso. Saldo atual: R$%.2f.\n", valor, c.Saldo)
	return nil
}

// Sacar retira um valor do saldo da conta
func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > c.Saldo {
		return &SaldoInsuficienteError{SaldoAtual: c.Saldo, ValorSaque: valor}
	}
	if valor <= 0 {
		return errors.New("O valor do saque deve ser positivo")
	}
	c.Saldo -= valor
	fmt.Printf("Saque de R$%.2f realizado com sucesso. Saldo atual: R$%.2f.\n", valor, c.Saldo)
	return nil
}

// ObterSaldo retorna o saldo atual da conta
func (c *ContaBancaria) ObterSaldo() float64 {
	return c.Saldo
}

// Função principal para teste
func main() {
	conta := ContaBancaria{Titular: "Carlos", Saldo: 1000}

	// Tentativa de saque que gera um erro
	err := conta.Sacar(1200)
	if err != nil {
		fmt.Println(err)
	}

	// Realizando um depósito
	err = conta.Depositar(500)
	if err != nil {
		fmt.Println(err)
	}

	// Realizando um saque válido
	err = conta.Sacar(200)
	if err != nil {
		fmt.Println(err)
	}
}
