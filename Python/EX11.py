from abc import ABC, abstractmethod

# Classe abstrata Funcionario
class Funcionario(ABC):
    @abstractmethod
    def calcularSalario(self):
        pass


# Classe derivada FuncionarioHorista
class FuncionarioHorista(Funcionario):
    def __init__(self, nome, horas_trabalhadas, valor_hora):
        self.nome = nome
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_hora = valor_hora

    def calcularSalario(self):
        return self.horas_trabalhadas * self.valor_hora


# Classe derivada FuncionarioAssalariado
class FuncionarioAssalariado(Funcionario):
    def __init__(self, nome, salario_mensal):
        self.nome = nome
        self.salario_mensal = salario_mensal

    def calcularSalario(self):
        return self.salario_mensal


# Exemplo de uso
if __name__ == "__main__":
    # Criando um funcionário horista
    funcionario_horista = FuncionarioHorista("Carlos", 160, 15)
    print(f"Salário de {funcionario_horista.nome}: R${funcionario_horista.calcularSalario():.2f}")

    # Criando um funcionário assalariado
    funcionario_assalariado = FuncionarioAssalariado("Melissa", 3000)
    print(f"Salário de {funcionario_assalariado.nome}: R${funcionario_assalariado.calcularSalario():.2f}")
