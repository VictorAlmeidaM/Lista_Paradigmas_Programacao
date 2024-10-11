class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, other):
        if isinstance(other, Produto):
            # Retorna um novo Produto com a soma dos preços
            return Produto(f"{self.nome} + {other.nome}", self.preco + other.preco)
        return NotImplemented  # Se 'other' não for um Produto, retorna NotImplemented

    def __str__(self):
        return f"Produto: {self.nome}, Preço: R${self.preco:.2f}"


# Exemplo de uso
if __name__ == "__main__":
    produto1 = Produto("Produto A", 50.00)
    produto2 = Produto("Produto B", 30.00)

    # Somando os produtos
    produto_total = produto1 + produto2

    # Exibindo o resultado
    print(produto_total)  # Saída: Produto: Produto A + Produto B, Preço: R$80.00
