class Calculadora {

    // Método somar para dois números
    public int somar(int num1, int num2) {
        return num1 + num2;
    }

    // Método somar para três números
    public int somar(int num1, int num2, int num3) {
        return num1 + num2 + num3;
    }

    public static void main(String[] args) {
        // Criando um objeto da Calculadora
        Calculadora calc = new Calculadora();

        // Soma de dois números
        int resultado1 = calc.somar(10, 20);
        System.out.println("Soma de 10 + 20: " + resultado1);

        // Soma de três números
        int resultado2 = calc.somar(10, 20, 30);
        System.out.println("Soma de 10 + 20 + 30: " + resultado2);
    }
}


