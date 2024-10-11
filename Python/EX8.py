# Classe Empregado
class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

    def exibir_informacoes(self):
        print(f"Nome: {self.nome}, Cargo: {self.cargo}, Salário: R${self.salario:.2f}")


# Classe Empresa
class Empresa:
    def __init__(self, nome_empresa):
        self.nome_empresa = nome_empresa
        self.empregados = []  # Lista para armazenar os empregados da empresa

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def listar_empregados(self):
        print(f"Lista de empregados da empresa {self.nome_empresa}:")
        for empregado in self.empregados:
            empregado.exibir_informacoes()

    def calcular_folha_pagamento(self):
        total_salarios = sum(empregado.salario for empregado in self.empregados)
        print(f"Folha de pagamento total: R${total_salarios:.2f}")


# Exemplo de uso
if __name__ == "__main__":
    # Criando empregados
    emp1 = Empregado("Joaquim", "Desenvolvedor", 5000)
    emp2 = Empregado("Marisa", "Gerente", 7500)
    emp3 = Empregado("Pedro", "Designer", 4000)

    # Criando uma empresa e adicionando empregados
    empresa = Empresa("Tech Solutions")
    empresa.adicionar_empregado(emp1)
    empresa.adicionar_empregado(emp2)
    empresa.adicionar_empregado(emp3)

    # Exibindo a lista de empregados e calculando a folha de pagamento
    empresa.listar_empregados()
    empresa.calcular_folha_pagamento()
