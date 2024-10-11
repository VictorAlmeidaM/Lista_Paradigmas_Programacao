class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano

    def exibir_detalhes(self):
        print(f'Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}')

# Instanciando três objetos da classe Carro
carro1 = Carro("Toyota", "Yaris", 2020)
carro2 = Carro("Honda", "Accord", 2019)
carro3 = Carro("Ford", "Edge", 2018)

# Exibindo os detalhes de cada carro
carro1.exibir_detalhes()
carro2.exibir_detalhes()
carro3.exibir_detalhes()
