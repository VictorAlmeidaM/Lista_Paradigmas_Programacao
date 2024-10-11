class Matematica {

    // Método estático para calcular o fatorial
    public static long fatorial(int n) {
        if (n < 0) {
            throw new IllegalArgumentException("O fatorial não está definido para números negativos.");
        } else if (n == 0 || n == 1) {
            return 1;
        } else {
            long resultado = 1;
            for (int i = 2; i <= n; i++) {
                resultado *= i;
            }
            return resultado;
        }
    }

    // Método principal para teste
    public static void main(String[] args) {
        int numero = 5;
        long resultado = Matematica.fatorial(numero);
        System.out.printf("O fatorial de %d é: %d%n", numero, resultado);
    }
}