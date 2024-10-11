class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0  # Velocidade inicial

    def exibir_detalhes(self):
        print(f'Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}')

    def acelerar(self, incremento):
        self.velocidade += incremento
        print(f'{self.marca} {self.modelo} acelerou. Velocidade atual: {self.velocidade} km/h')

    def frear(self, decremento):
        self.velocidade -= decremento
        if self.velocidade < 0:
            self.velocidade = 0  # Garantindo que a velocidade não seja negativa
        print(f'{self.marca} {self.modelo} freou. Velocidade atual: {self.velocidade} km/h')

    def exibir_velocidade(self):
        print(f'Velocidade atual do {self.marca} {self.modelo}: {self.velocidade} km/h')

# Instanciando três objetos da classe Carro
carro1 = Carro("Toyota", "Yaris", 2020)
carro2 = Carro("Honda", "Accord", 2019)
carro3 = Carro("Ford", "Edge", 2018)

# Exibindo detalhes e controlando a velocidade dos carros
carro1.exibir_detalhes()
carro1.acelerar(30)
carro1.frear(10)
carro1.exibir_velocidade()

carro2.exibir_detalhes()
carro2.acelerar(50)
carro2.frear(20)
carro2.exibir_velocidade()

carro3.exibir_detalhes()
carro3.acelerar(40)
carro3.frear(50)  # Tentativa de frear mais do que a velocidade
carro3.exibir_velocidade()
