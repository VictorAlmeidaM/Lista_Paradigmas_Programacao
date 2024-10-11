class ContaBancaria{
    private double saldo; // Atributo privado, protegido por encapsulamento
    private String titular;

    // Construtor
    public ContaBancaria(String titular, double saldoInicial) {
        this.titular = titular;
        this.saldo = saldoInicial;
    }

    // Sobrecarga do construtor para saldo inicial opcional
    public ContaBancaria(String titular) {
        this(titular, 0); // Chama o construtor principal com saldo inicial = 0
    }

    // Método para depositar dinheiro
    public void depositar(double valor) {
        if (valor > 0) {
            saldo += valor;
            System.out.println("Depósito de R$" + valor + " realizado com sucesso. Saldo atual: R$" + saldo);
        } else {
            System.out.println("Valor de depósito inválido.");
        }
    }

    // Método para sacar dinheiro
    public void sacar(double valor) {
        if (valor > 0 && valor <= saldo) {
            saldo -= valor;
            System.out.println("Saque de R$" + valor + " realizado com sucesso. Saldo atual: R$" + saldo);
        } else {
            System.out.println("Saldo insuficiente ou valor de saque inválido.");
        }
    }

    // Método para exibir o saldo
    public void exibirSaldo() {
        System.out.println("Saldo atual de " + titular + ": R$" + saldo);
    }

    // Método principal para testar a classe
    public static void main(String[] args) {
        // Exemplo de uso:
        ContaBancaria conta1 = new ContaBancaria("José", 1000);
        conta1.exibirSaldo();
        conta1.depositar(500);
        conta1.sacar(300);
        conta1.exibirSaldo();

        ContaBancaria conta2 = new ContaBancaria("Marina", 200);
        conta2.exibirSaldo();
        conta2.sacar(500);  // Tentativa de saque maior que o saldo
        conta2.depositar(100);
        conta2.sacar(50);
        conta2.exibirSaldo();
    }
}