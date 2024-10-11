class Carro {
    private String marca;
    private String modelo;
    private int ano;

    // Construtor
    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
    }

    // Método para exibir os detalhes
    public void exibirDetalhes() {
        System.out.println("Marca: " + this.marca + ", Modelo: " + this.modelo + ", Ano: " + this.ano);
    }

    // Método principal
    public static void main(String[] args) {
        // Instanciando três objetos da classe Carro
        Carro carro1 = new Carro("Toyota", "Yaris", 2020);
        Carro carro2 = new Carro("Honda", "Accord", 2019);
        Carro carro3 = new Carro("Ford", "Edge", 2018);

        // Exibindo os detalhes de cada carro
        carro1.exibirDetalhes();
        carro2.exibirDetalhes();
        carro3.exibirDetalhes();
    }
}
