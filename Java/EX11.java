
// Classe abstrata Funcionario
abstract class Funcionario {
    protected String nome;

    // Método abstrato para calcular salário
    public abstract double calcularSalario();

    // Método para obter o nome do funcionário
    public String getNome() {
        return nome;
    }
}

// Classe derivada FuncionarioHorista
class FuncionarioHorista extends Funcionario {
    private int horasTrabalhadas;
    private double valorHora;

    // Construtor
    public FuncionarioHorista(String nome, int horasTrabalhadas, double valorHora) {
        this.nome = nome;
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorHora = valorHora;
    }

    // Implementação do método calcularSalario
    @Override
    public double calcularSalario() {
        return horasTrabalhadas * valorHora;
    }
}

// Classe derivada FuncionarioAssalariado
class FuncionarioAssalariado extends Funcionario {
    private double salarioMensal;

    // Construtor
    public FuncionarioAssalariado(String nome, double salarioMensal) {
        this.nome = nome;
        this.salarioMensal = salarioMensal;
    }

    // Implementação do método calcularSalario
    @Override
    public double calcularSalario() {
        return salarioMensal;
    }
}

// Classe principal para teste
class Main6 {
    public static void main(String[] args) {
        // Criando um funcionário horista
        Funcionario funcionarioHorista = new FuncionarioHorista("Carlos", 203, 15);
        System.out.printf("Salário de %s: R$%.2f%n", funcionarioHorista.getNome(), funcionarioHorista.calcularSalario());

        // Criando um funcionário assalariado
        Funcionario funcionarioAssalariado = new FuncionarioAssalariado("Melissa", 3000);
        System.out.printf("Salário de %s: R$%.2f%n", funcionarioAssalariado.getNome(), funcionarioAssalariado.calcularSalario());
    }
}
