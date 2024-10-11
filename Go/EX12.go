package main

import (
	"fmt"
)

// Estrutura Produto
type Produto struct {
	nome  string
	preco float64
}

// Construtor para criar um novo produto
func NewProduto(nome string, preco float64) Produto {
	return Produto{nome: nome, preco: preco}
}

// Método para somar dois produtos
func (p Produto) Somar(outro Produto) Produto {
	// Retorna um novo Produto com a soma dos preços
	return Produto{
		nome:  p.nome + " + " + outro.nome,
		preco: p.preco + outro.preco,
	}
}

// Método para exibir informações do produto
func (p Produto) String() string {
	return fmt.Sprintf("Produto: %s, Preço: R$%.2f", p.nome, p.preco)
}

// Função principal
func main() {
	produto1 := NewProduto("Produto A", 80.00)
	produto2 := NewProduto("Produto B", 30.00)

	// Somando os produtos
	produtoTotal := produto1.Somar(produto2)

	// Exibindo o resultado
	fmt.Println(produtoTotal.String()) // Saída: Produto: Produto A + Produto B, Preço: R$110.00
}
