from abc import ABC, abstractmethod

# Definindo a interface (classe abstrata) Imprimivel
class Imprimivel(ABC):
    @abstractmethod
    def imprimir(self):
        pass

# Classe Relatorio que implementa a interface Imprimivel
class Relatorio(Imprimivel):
    def __init__(self, titulo, conteudo):
        self.titulo = titulo
        self.conteudo = conteudo

    # Implementação do método imprimir
    def imprimir(self):
        print(f"Relatório: {self.titulo}")
        print(f"Conteúdo: {self.conteudo}")

# Classe Contrato que implementa a interface Imprimivel
class Contrato(Imprimivel):
    def __init__(self, partes_envolvidas, termos):
        self.partes_envolvidas = partes_envolvidas
        self.termos = termos

    # Implementação do método imprimir
    def imprimir(self):
        print(f"Contrato entre: {self.partes_envolvidas}")
        print(f"Termos do contrato: {self.termos}")

# Exemplo de uso
if __name__ == "__main__":
    # Criando um relatório e um contrato
    relatorio = Relatorio("Relatório de Vendas", "Vendas aumentaram 20% este mês.")
    contrato = Contrato("Empresa A e Empresa B", "Termos de fornecimento por 12 meses.")

    # Imprimindo o relatório e o contrato
    relatorio.imprimir()
    contrato.imprimir()
