import java.util.ArrayList;
import java.util.List;

class Professor {
    private String nome;
    private List<Escola> escolas; // Lista para armazenar as escolas que o professor leciona

    public Professor(String nome) {
        this.nome = nome;
        this.escolas = new ArrayList<>();
    }

    public void adicionarEscola(Escola escola) {
        if (!escolas.contains(escola)) {
            escolas.add(escola);
            escola.adicionarProfessor(this); // Adiciona o professor à escola também
        }
    }

    @Override
    public String toString() {
        return nome;
    }
}

class Escola {
    private String nome;
    private List<Professor> professores; // Lista para armazenar os professores da escola

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public void adicionarProfessor(Professor professor) {
        if (!professores.contains(professor)) {
            professores.add(professor);
        }
    }

    @Override
    public String toString() {
        return nome;
    }

    public List<String> listarProfessores() {
        List<String> nomesProfessores = new ArrayList<>();
        for (Professor professor : professores) {
            nomesProfessores.add(professor.toString());
        }
        return nomesProfessores;
    }
}

class Main3 {
    public static void main(String[] args) {
        Escola escola1 = new Escola("Escola A");
        Escola escola2 = new Escola("Escola B");

        Professor professor1 = new Professor("Roberto");
        Professor professor2 = new Professor("Fernanda");

        // Associando professores às escolas
        professor1.adicionarEscola(escola1);
        professor1.adicionarEscola(escola2);
        professor2.adicionarEscola(escola1);

        // Exibindo os professores de cada escola
        System.out.println("Professores da " + escola1 + ": " + escola1.listarProfessores());
        System.out.println("Professores da " + escola2 + ": " + escola2.listarProfessores());
    }
}