class Calculadora:
    # Método somar para dois parâmetros
    def somar_dois(self, num1, num2):
        return num1 + num2

    # Método somar para três parâmetros
    def somar_tres(self, num1, num2, num3):
        return num1 + num2 + num3


# Exemplo de uso
if __name__ == "__main__":
    calc = Calculadora()

    # Soma de dois números
    resultado1 = calc.somar_dois(10, 20)
    print(f"Soma de 10 + 20: {resultado1}")

    # Soma de três números
    resultado2 = calc.somar_tres(10, 20, 30)
    print(f"Soma de 10 + 20 + 30: {resultado2}")
