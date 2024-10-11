// Definição da interface Imprimivel
interface Imprimivel {
    void imprimir();
}

// Classe Relatorio que implementa a interface Imprimivel
class Relatorio implements Imprimivel {
    private String titulo;
    private String conteudo;

    // Construtor
    public Relatorio(String titulo, String conteudo) {
        this.titulo = titulo;
        this.conteudo = conteudo;
    }

    // Implementação do método imprimir
    @Override
    public void imprimir() {
        System.out.println("Relatório: " + titulo);
        System.out.println("Conteúdo: " + conteudo);
    }
}

// Classe Contrato que implementa a interface Imprimivel
class Contrato implements Imprimivel {
    private String partesEnvolvidas;
    private String termos;

    // Construtor
    public Contrato(String partesEnvolvidas, String termos) {
        this.partesEnvolvidas = partesEnvolvidas;
        this.termos = termos;
    }

    // Implementação do método imprimir
    @Override
    public void imprimir() {
        System.out.println("Contrato entre: " + partesEnvolvidas);
        System.out.println("Termos do contrato: " + termos);
    }
}

// Classe principal para testar Relatorio e Contrato
class Main5 {
    public static void main(String[] args) {
        // Criando um relatório e um contrato
        Imprimivel relatorio = new Relatorio("Relatório de Vendas", "Vendas aumentaram 20% este mês.");
        Imprimivel contrato = new Contrato("Empresa A e Empresa B", "Termos de fornecimento por 12 meses.");

        // Imprimindo o relatório e o contrato
        relatorio.imprimir();
        contrato.imprimir();
    }
}
