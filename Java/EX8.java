import java.util.ArrayList;
import java.util.List;

// Classe Empregado
class Empregado {
    private String nome;
    private String cargo;
    private double salario;

    // Construtor
    public Empregado(String nome, String cargo, double salario) {
        this.nome = nome;
        this.cargo = cargo;
        this.salario = salario;
    }

    // Método para exibir informações do empregado
    public void exibirInformacoes() {
        System.out.printf("Nome: %s, Cargo: %s, Salário: R$%.2f%n", nome, cargo, salario);
    }

    // Getter para o salário
    public double getSalario() {
        return salario;
    }
}

// Classe Empresa
class Empresa {
    private String nomeEmpresa;
    private List<Empregado> empregados; // Lista para armazenar os empregados da empresa

    // Construtor
    public Empresa(String nomeEmpresa) {
        this.nomeEmpresa = nomeEmpresa;
        this.empregados = new ArrayList<>();
    }

    // Método para adicionar um empregado
    public void adicionarEmpregado(Empregado empregado) {
        empregados.add(empregado);
    }

    // Método para listar os empregados da empresa
    public void listarEmpregados() {
        System.out.println("Lista de empregados da empresa " + nomeEmpresa + ":");
        for (Empregado empregado : empregados) {
            empregado.exibirInformacoes();
        }
    }

    // Método para calcular a folha de pagamento
    public void calcularFolhaPagamento() {
        double totalSalarios = 0;
        for (Empregado empregado : empregados) {
            totalSalarios += empregado.getSalario();
        }
        System.out.printf("Folha de pagamento total: R$%.2f%n", totalSalarios);
    }
}

// Classe principal com o método main
class Main4 {
    public static void main(String[] args) {
        // Criando empregados
        Empregado emp1 = new Empregado("Joaquim", "Desenvolvedor", 5000);
        Empregado emp2 = new Empregado("Marisa", "Gerente", 7500);
        Empregado emp3 = new Empregado("Pedro", "Designer", 4000);

        // Criando uma empresa e adicionando empregados
        Empresa empresa = new Empresa("Tech Solutions");
        empresa.adicionarEmpregado(emp1);
        empresa.adicionarEmpregado(emp2);
        empresa.adicionarEmpregado(emp3);

        // Exibindo a lista de empregados e calculando a folha de pagamento
        empresa.listarEmpregados();
        empresa.calcularFolhaPagamento();
    }
}