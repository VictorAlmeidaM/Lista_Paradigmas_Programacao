import java.util.ArrayList;
import java.util.List;

// Classe base Animal
abstract class Animal2 {
    // Método abstrato para ser implementado pelas subclasses
    public abstract String som();
}

// Classe derivada Cachorro
class Cachorro2 extends Animal {
    @Override
    public String som() {
        return "Au Au";
    }
}

// Classe derivada Gato
class Gato2 extends Animal {
    @Override
    public String som() {
        return "Miau";
    }
}

// Classe principal
class Main1 {
    // Função que recebe uma lista de objetos Animal
    public static void emitirSons(List<Animal> animais) {
        for (Animal animal : animais) {
            System.out.println(animal.som());
        }
    }

    public static void main(String[] args) {
        // Criando instâncias das classes
        Cachorro cachorro1 = new Cachorro();
        Gato gato1 = new Gato();
        Cachorro cachorro2 = new Cachorro();

        // Criando uma lista de animais
        List<Animal> animais = new ArrayList<>();
        animais.add(cachorro1);
        animais.add(gato1);
        animais.add(cachorro2);

        // Chamando a função que emite os sons
        emitirSons(animais);
    }
}