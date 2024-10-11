# Classe base Animal
class Animal:
    def som(self):
        raise NotImplementedError("Subclasse deve implementar o método som")

# Classe derivada Cachorro
class Cachorro(Animal):
    def som(self):
        return "Au Au"

# Classe derivada Gato
class Gato(Animal):
    def som(self):
        return "Miau"

# Função que recebe uma lista de objetos Animal
def emitir_sons(animais):
    for animal in animais:
        print(animal.som())

# Exemplo de uso
def main():
    # Criando instâncias das classes
    cachorro1 = Cachorro()
    gato1 = Gato()
    cachorro2 = Cachorro()
    
    # Criando uma lista de animais
    animais = [cachorro1, gato1, cachorro2]

    # Chamando a função que emite os sons
    emitir_sons(animais)

if __name__ == "__main__":
    main()
