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

# Exemplo de uso
def main():
    # Criando instâncias das classes
    cachorro = Cachorro()
    gato = Gato()

    # Exibindo os sons dos animais
    print(f"Cachorro: {cachorro.som()}")
    print(f"Gato: {gato.som()}")

if __name__ == "__main__":
    main()
