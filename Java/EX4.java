// Classe base Animal
abstract class Animal {
    // Método abstrato para ser implementado pelas subclasses
    public abstract String som();
}

// Classe derivada Cachorro
class Cachorro extends Animal {
    @Override
    public String som() {
        return "Au Au";
    }
}

// Classe derivada Gato
class Gato extends Animal {
    @Override
    public String som() {
        return "Miau";
    }
}

// Classe principal
class Main {
    public static void main(String[] args) {
        // Criando instâncias das classes
        Animal cachorro = new Cachorro();
        Animal gato = new Gato();

        // Exibindo os sons dos animais
        System.out.println("Cachorro: " + cachorro.som());
        System.out.println("Gato: " + gato.som());
    }
}