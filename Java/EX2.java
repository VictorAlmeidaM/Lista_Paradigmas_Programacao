class Carro3 {
    private String marca;
    private String modelo;
    private int ano;
    private int velocidade;

    // Construtor
    public Carro3(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0; // Velocidade inicial
    }

    // Método para exibir os detalhes
    public void exibirDetalhes() {
        System.out.println("Marca: " + this.marca + ", Modelo: " + this.modelo + ", Ano: " + this.ano);
    }

    // Método para acelerar
    public void acelerar(int incremento) {
        this.velocidade += incremento;
        System.out.println(this.marca + " " + this.modelo + " acelerou. Velocidade atual: " + this.velocidade + " km/h");
    }

    // Método para frear
    public void frear(int decremento) {
        this.velocidade -= decremento;
        if (this.velocidade < 0) {
            this.velocidade = 0; // Garantindo que a velocidade não seja negativa
        }
        System.out.println(this.marca + " " + this.modelo + " freou. Velocidade atual: " + this.velocidade + " km/h");
    }

    // Método para exibir a velocidade atual
    public void exibirVelocidade() {
        System.out.println("Velocidade atual do " + this.marca + " " + this.modelo + ": " + this.velocidade + " km/h");
    }

    // Método principal para testar a classe
    public static void main(String[] args) {
        // Instanciando três objetos da classe Carro
        Carro3 carro1 = new Carro3("Toyota", "Yaris", 2020);
        Carro3 carro2 = new Carro3("Honda", "Accord", 2019);
        Carro3 carro3 = new Carro3("Ford", "Edge", 2018);

        // Exibindo detalhes e controlando a velocidade dos carros
        carro1.exibirDetalhes();
        carro1.acelerar(30);
        carro1.frear(10);
        carro1.exibirVelocidade();

        carro2.exibirDetalhes();
        carro2.acelerar(50);
        carro2.frear(20);
        carro2.exibirVelocidade();

        carro3.exibirDetalhes();
        carro3.acelerar(40);
        carro3.frear(50); // Tentativa de frear mais do que a velocidade
        carro3.exibirVelocidade();
    }
}