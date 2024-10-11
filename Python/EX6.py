class Motor:
    def __init__(self, potencia, tipo_combustivel):
        self.potencia = potencia  # em cavalos de potência (cv)
        self.tipo_combustivel = tipo_combustivel  # gasolina, etanol, etc.

    def exibir_informacoes(self):
        print(f"Motor: {self.potencia} cv, Tipo de Combustível: {self.tipo_combustivel}")


class Carro:
    def __init__(self, marca, modelo, ano, motor):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.motor = motor  # Atributo que é um objeto da classe Motor

    def exibir_detalhes(self):
        print(f"Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}")
        self.motor.exibir_informacoes()  # Exibe informações do motor associado


if __name__ == "__main__":
    # Criando um objeto Motor
    motor1 = Motor(150, "Gasolina")

    # Criando um objeto Carro associado ao Motor
    carro1 = Carro("Toyota", "Yaris", 2023, motor1)

    # Exibindo detalhes do carro e do motor
    carro1.exibir_detalhes()
