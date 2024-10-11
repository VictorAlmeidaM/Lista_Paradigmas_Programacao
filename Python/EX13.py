class Matematica:
    @staticmethod
    def fatorial(n):
        if n < 0:
            raise ValueError("O fatorial não está definido para números negativos.")
        elif n == 0 or n == 1:
            return 1
        else:
            resultado = 1
            for i in range(2, n + 1):
                resultado *= i
            return resultado


# Exemplo de uso
if __name__ == "__main__":
    numero = 6
    resultado = Matematica.fatorial(numero)
    print(f"O fatorial de {numero} é: {resultado}")
