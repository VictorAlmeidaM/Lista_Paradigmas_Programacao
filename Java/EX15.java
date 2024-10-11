// Classe de exceção personalizada
class SaldoInsuficienteException extends Exception {
    private double saldoAtual;
    private double valorSaque;

    public SaldoInsuficienteException(double saldoAtual, double valorSaque) {
        super(String.format("Saldo insuficiente: R$%.2f para o saque de R$%.2f.", saldoAtual, valorSaque));
        this.saldoAtual = saldoAtual;
        this.valorSaque = valorSaque;
    }

    public double getSaldoAtual() {
        return saldoAtual;
    }

    public double getValorSaque() {
        return valorSaque;
    }
}

// Classe ContaBancaria
class ContaBancaria1 {
    private String titular;
    private double saldo;

    public ContaBancaria1(String titular, double saldoInicial) {
        this.titular = titular;
        this.saldo = saldoInicial;
    }

    public void depositar(double valor) {
        if (valor > 0) {
            saldo += valor;
            System.out.printf("Depósito de R$%.2f realizado com sucesso. Saldo atual: R$%.2f.%n", valor, saldo);
        } else {
            System.out.println("O valor do depósito deve ser positivo.");
        }
    }

    public void sacar(double valor) throws SaldoInsuficienteException {
        if (valor > saldo) {
            throw new SaldoInsuficienteException(saldo, valor);
        }
        if (valor > 0) {
            saldo -= valor;
            System.out.printf("Saque de R$%.2f realizado com sucesso. Saldo atual: R$%.2f.%n", valor, saldo);
        } else {
            System.out.println("O valor do saque deve ser positivo.");
        }
    }

    public double obterSaldo() {
        return saldo;
    }
}

// Classe principal para teste
class Main8 {
    public static void main(String[] args) {
        ContaBancaria1 conta = new ContaBancaria1("Carlos", 1000);

        try {
            conta.sacar(1200);  // Tentativa de saque que gera exceção
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }

        conta.depositar(500);
        System.out.printf("Saldo atual: R$%.2f%n", conta.obterSaldo());

        try {
            conta.sacar(200);  // Saque válido
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }
    }
}


