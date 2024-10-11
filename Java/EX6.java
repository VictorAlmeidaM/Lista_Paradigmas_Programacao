// Classe Motor
class Motor {
    private int potencia; // em cavalos de potência (cv)
    private String tipoCombustivel; // gasolina, etanol, etc.

    // Construtor
    public Motor(int potencia, String tipoCombustivel) {
        this.potencia = potencia;
        this.tipoCombustivel = tipoCombustivel;
    }

    // Método para exibir informações do motor
    public void exibirInformacoes() {
        System.out.println("Motor: " + potencia + " cv, Tipo de Combustível: " + tipoCombustivel);
    }
}

// Classe Carro
class Carro2 {
    private String marca;
    private String modelo;
    private int ano;
    private Motor motor; // Atributo que é um objeto da classe Motor

    // Construtor
    public Carro2(String marca, String modelo, int ano, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.motor = motor;
    }

    // Método para exibir detalhes do carro
    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano);
        motor.exibirInformacoes(); // Exibe informações do motor associado
    }
}

// Classe principal
class Main2 {
    public static void main(String[] args) {
        // Criando um objeto Motor
        Motor motor1 = new Motor(150, "Gasolina");

        // Criando um objeto Carro associado ao Motor
        Carro2 carro1 = new Carro2("Toyota", "Yaris", 2023, motor1);

        // Exibindo detalhes do carro e do motor
        carro1.exibirDetalhes();
    }
}