class ContaBancaria:
    def __init__(self, titular, saldo_inicial=0):
        self.__saldo = saldo_inicial  # Atributo privado, protegido por encapsulamento
        self.titular = titular

    # Método para depositar dinheiro
    def depositar(self, valor):
        if valor > 0:
            self.__saldo += valor
            print(f"Depósito de R${valor} realizado com sucesso. Saldo atual: R${self.__saldo}")
        else:
            print("Valor de depósito inválido.")

    # Método para sacar dinheiro
    def sacar(self, valor):
        if 0 < valor <= self.__saldo:
            self.__saldo -= valor
            print(f"Saque de R${valor} realizado com sucesso. Saldo atual: R${self.__saldo}")
        else:
            print("Saldo insuficiente ou valor de saque inválido.")

    # Método para exibir o saldo
    def exibir_saldo(self):
        print(f"Saldo atual de {self.titular}: R${self.__saldo}")

# Exemplo de uso:
conta1 = ContaBancaria("José", 1000)
conta1.exibir_saldo()
conta1.depositar(500)
conta1.sacar(300)
conta1.exibir_saldo()

conta2 = ContaBancaria("Marina", 200)
conta2.exibir_saldo()
conta2.sacar(500)  # Tentativa de saque maior que o saldo
conta2.depositar(100)
conta2.sacar(50)
conta2.exibir_saldo()
