class SaldoInsuficienteException(Exception):
    """Exceção lançada quando o saldo é insuficiente para um saque."""
    def __init__(self, saldo_atual, valor_saque):
        self.saldo_atual = saldo_atual
        self.valor_saque = valor_saque
        super().__init__(f"Saldo insuficiente: R${self.saldo_atual:.2f} para o saque de R${self.valor_saque:.2f}.")

class ContaBancaria:
    def __init__(self, titular, saldo_inicial=0):
        self.titular = titular
        self.saldo = saldo_inicial

    def depositar(self, valor):
        if valor > 0:
            self.saldo += valor
            print(f"Deposito de R${valor:.2f} realizado com sucesso. Saldo atual: R${self.saldo:.2f}.")
        else:
            print("O valor do depósito deve ser positivo.")

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException(self.saldo, valor)
        if valor > 0:
            self.saldo -= valor
            print(f"Saque de R${valor:.2f} realizado com sucesso. Saldo atual: R${self.saldo:.2f}.")
        else:
            print("O valor do saque deve ser positivo.")

    def obter_saldo(self):
        return self.saldo

# Exemplo de uso
if __name__ == "__main__":
    conta = ContaBancaria("Carlos", 1000)

    try:
        conta.sacar(1200)  # Tentativa de saque que gera exceção
    except SaldoInsuficienteException as e:
        print(e)

    conta.depositar(500)
    print(f"Saldo atual: R${conta.obter_saldo():.2f}")

    try:
        conta.sacar(200)  # Saque válido
    except SaldoInsuficienteException as e:
        print(e)
