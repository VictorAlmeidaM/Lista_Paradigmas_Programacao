package main

import "fmt"

// Definição da interface Imprimivel
type Imprimivel interface {
	imprimir()
}

// Estrutura Relatorio que implementa a interface Imprimivel
type Relatorio struct {
	titulo   string
	conteudo string
}

// Construtor de Relatorio
func novoRelatorio(titulo, conteudo string) *Relatorio {
	return &Relatorio{
		titulo:   titulo,
		conteudo: conteudo,
	}
}

// Implementação do método imprimir para Relatorio
func (r *Relatorio) imprimir() {
	fmt.Printf("Relatório: %s\n", r.titulo)
	fmt.Printf("Conteúdo: %s\n", r.conteudo)
}

// Estrutura Contrato que implementa a interface Imprimivel
type Contrato struct {
	partesEnvolvidas string
	termos           string
}

// Construtor de Contrato
func novoContrato(partesEnvolvidas, termos string) *Contrato {
	return &Contrato{
		partesEnvolvidas: partesEnvolvidas,
		termos:           termos,
	}
}

// Implementação do método imprimir para Contrato
func (c *Contrato) imprimir() {
	fmt.Printf("Contrato entre: %s\n", c.partesEnvolvidas)
	fmt.Printf("Termos do contrato: %s\n", c.termos)
}

// Função principal para testar Relatorio e Contrato
func main() {
	// Criando um relatório e um contrato
	var relatorio Imprimivel = novoRelatorio("Relatório de Vendas", "Vendas aumentaram 20% este mês.")
	var contrato Imprimivel = novoContrato("Empresa A e Empresa B", "Termos de fornecimento por 12 meses.")

	// Imprimindo o relatório e o contrato
	relatorio.imprimir()
	contrato.imprimir()
}
