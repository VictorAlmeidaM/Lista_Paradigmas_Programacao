import java.util.Properties;

class Configuracao {
    // Atributo estático para armazenar a instância única da classe
    private static Configuracao instancia;

    // Atributo para armazenar as configurações
    private final Properties configuracao;

    // Construtor privado para evitar a criação de instâncias externas
    private Configuracao() {
        configuracao = new Properties(); // Inicializa as configurações
    }

    // Método estático para obter a instância única da classe
    public static Configuracao getInstancia() {
        if (instancia == null) { // Verifica se a instância já foi criada
            instancia = new Configuracao();
        }
        return instancia;
    }

    // Método para definir uma configuração
    public void setConfig(String chave, String valor) {
        configuracao.setProperty(chave, valor);
    }

    // Método para obter uma configuração
    public String getConfig(String chave) {
        return configuracao.getProperty(chave);
    }
}

// Exemplo de uso da classe Configuracao
class Main7 {
    public static void main(String[] args) {
        // Obtendo a instância única da classe Configuracao
        Configuracao config1 = Configuracao.getInstancia();
        Configuracao config2 = Configuracao.getInstancia();

        // Definindo configurações
        config1.setConfig("url", "https://example.com");
        config1.setConfig("timeout", "30");

        // Obtendo configurações usando a segunda instância
        System.out.println("URL: " + config2.getConfig("url")); // Saída: URL: https://example.com
        System.out.println("Timeout: " + config2.getConfig("timeout")); // Saída: Timeout: 30

        // Verificando se ambas as instâncias são a mesma
        System.out.println("config1 é config2? " + (config1 == config2)); // Saída: config1 é config2? true
    }
}


