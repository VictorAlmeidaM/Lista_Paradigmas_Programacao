// Classe Produto
class Produto {
    private String nome;
    private double preco;

    // Construtor
    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    // Método para somar dois produtos
    public Produto somar(Produto outro) {
        // Retorna um novo Produto com a soma dos preços
        return new Produto(this.nome + " + " + outro.nome, this.preco + outro.preco);
    }

    // Método para exibir informações do produto
    @Override
    public String toString() {
        return "Produto: " + nome + ", Preço: R$" + String.format("%.2f", preco);
    }

    // Método principal para testar a funcionalidade
    public static void main(String[] args) {
        Produto produto1 = new Produto("Produto A", 50.00);
        Produto produto2 = new Produto("Produto B", 30.00);

        // Somando os produtos
        Produto produtoTotal = produto1.somar(produto2);

        // Exibindo o resultado
        System.out.println(produtoTotal);  // Saída: Produto: Produto A + Produto B, Preço: R$80.00
    }
}



